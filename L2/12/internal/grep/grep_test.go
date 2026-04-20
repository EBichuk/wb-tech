package grep

import (
	"bytes"
	"grep/internal/config"
	"os"
	"reflect"
	"testing"
)

func TestGetContext(t *testing.T) {
	tests := []struct {
		name                               string
		context, after, before, start, end int
	}{
		{"after+before", 0, 1, 1, 1, 1},
		{"contex>after && contex>before", 2, 1, 1, 2, 2},
		{"contex<after && contex>before", 1, 2, 0, 1, 2},
		{"contex>after && contex<before", 1, 0, 2, 2, 1},
		{"contex==after && contex==before", 1, 1, 1, 1, 1},
		{"everything is 0", 0, 0, 0, 0, 0},
		{"non positive context", -1, 0, 0, 0, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotStart, gotEnd := GetContext(tt.context, tt.after, tt.before)

			if !reflect.DeepEqual(gotStart, tt.start) {
				t.Errorf("GetContext() = %v, want %v", gotStart, tt.start)
			}
			if !reflect.DeepEqual(gotEnd, tt.end) {
				t.Errorf("GetContext() = %v, want %v", gotEnd, tt.end)
			}
		})
	}
}

func RunWithStdin(t *testing.T, cfg *config.Config, input string) string {
	oldStdin := os.Stdin
	r, w, _ := os.Pipe()

	os.Stdin = r

	go func() {
		defer w.Close()
		w.WriteString(input)
	}()

	var dst bytes.Buffer
	grep := &Grep{
		cfg: cfg,
		dst: &dst,
	}

	grep.Run()

	os.Stdin = oldStdin
	r.Close()

	return dst.String()
}

func TestRun(t *testing.T) {
	tests := []struct {
		name     string
		cfg      *config.Config
		input    string
		expected string
	}{
		{
			name: "simple match",
			cfg: &config.Config{
				Pattern: "line",
			},
			input:    "this is 1 line\ngap\nthe third line\nGAP\n",
			expected: "this is 1 line\nthe third line\n",
		},
		{
			name: "simple match: only digits",
			cfg: &config.Config{
				Pattern: "^\\d+$",
			},
			input:    "123 43\n42\n4O2\n8a32\n505\n",
			expected: "42\n505\n",
		},
		{
			name: "no match",
			cfg: &config.Config{
				Pattern: "THE",
			},
			input:    "this is 1 line\ngap\nthe third line\nGAP\n",
			expected: "",
		},
		{
			name: "repeat string",
			cfg: &config.Config{
				Pattern: "gap",
			},
			input:    "this is 1 line\ngap\ngap\nthe third line\ngap\n",
			expected: "gap\ngap\ngap\n",
		},
		{
			name: "count only",
			cfg: &config.Config{
				Pattern: "line",
				Count:   true,
			},
			input:    "this is 1 line\ngap\nthe third line\nGAP\n",
			expected: "2\n",
		},
		{
			name: "count only with no match",
			cfg: &config.Config{
				Pattern: "spring",
				Count:   true,
			},
			input:    "this is 1 line\ngap\nthe third line\nGAP\n",
			expected: "0\n",
		},
		{
			name: "context",
			cfg: &config.Config{
				Pattern: "line",
				Context: 1,
			},
			input:    "this is 1 line\ngap\nthe third line\nGAP\nend\n",
			expected: "this is 1 line\ngap\nthe third line\nGAP\n",
		},
		{
			name: "only before context",
			cfg: &config.Config{
				Pattern:       "line",
				BeforeContext: 1,
			},
			input:    "this is 1 line\ngap\nthe third line\nGAP\nend\n",
			expected: "this is 1 line\ngap\nthe third line\n",
		},
		{
			name: "only after context",
			cfg: &config.Config{
				Pattern:      "gap",
				AfterContext: 1,
			},
			input:    "this is 1 line\ngap\nthe third line\nGAP\nend\n",
			expected: "gap\nthe third line\n",
		},
		{
			name: "ignore case",
			cfg: &config.Config{
				Pattern:    "gap",
				IgnoreCase: true,
			},
			input:    "this is 1 line\ngap\nthe third line\nGAP\nend\n",
			expected: "gap\nGAP\n",
		},
		{
			name: "invert match",
			cfg: &config.Config{
				Pattern: "line",
				Invert:  true,
			},
			input:    "this is 1 line\ngap\nthe third line\nGAP\nend\n",
			expected: "gap\nGAP\nend\n",
		},
		{
			name: "line number",
			cfg: &config.Config{
				Pattern: "^\\d+$",
				LineNum: true,
			},
			input:    "123 43\n42\n4O2\n8a32\n505\n",
			expected: "2:42\n5:505\n",
		},
		{
			name: "line numbers invert match and context",
			cfg: &config.Config{
				Pattern: "line",
				Invert:  true,
				LineNum: true,
				Context: 1,
			},
			input:    "this is 1 line\ngap\nthe third line\nGAP\nend\n",
			expected: "1-this is 1 line\n2:gap\n3-the third line\n4:GAP\n5:end\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := RunWithStdin(t, tt.cfg, tt.input)
			if got != tt.expected {
				t.Errorf("expected:\n%q\ngot:\n%q", tt.expected, got)
			}
		})
	}
}
