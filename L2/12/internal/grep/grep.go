package grep

import (
	"bufio"
	"fmt"
	"grep/internal/config"
	"io"
	"os"
	"regexp"
)

type Grep struct {
	cfg *config.Config
	dst io.Writer
}

func New(config *config.Config, dst io.Writer) *Grep {
	return &Grep{
		cfg: config,
		dst: dst,
	}
}

func (g *Grep) Parse() ([]string, error) {
	var srs io.Reader
	if g.cfg.FileName == "" {
		srs = os.Stdin
	} else {
		file, err := os.Open(g.cfg.FileName)
		if err != nil {
			return nil, err
		}
		srs = file
		defer file.Close()
	}

	scanner := bufio.NewScanner(srs)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func (g *Grep) Run() {
	lines, err := g.Parse()
	if err != nil {
		fmt.Fprint(g.dst, err.Error())
		return
	}

	re := GetPattern(g.cfg.Pattern, g.cfg.Fixed, g.cfg.IgnoreCase)

	matchLine := make([]bool, len(lines))
	printLines := make(map[int]struct{})
	var countMatched int

	start, end := GetContext(g.cfg.Context, g.cfg.AfterContext, g.cfg.BeforeContext)

	for i, line := range lines {
		if !g.cfg.Invert && re.MatchString(line) || g.cfg.Invert && !re.MatchString(line) {
			matchLine[i] = true
			countMatched++

			for j := i - start; j < i; j++ {
				if j >= 0 {
					printLines[j] = struct{}{}
				}
			}
			for j := i; j <= i+end; j++ {
				if j < len(lines) {
					printLines[j] = struct{}{}
				}
			}
		}
	}

	if g.cfg.Count {
		fmt.Fprintln(g.dst, countMatched)
		return
	}

	if !g.cfg.LineNum {
		for i := range matchLine {
			if _, ok := printLines[i]; ok {
				fmt.Fprintln(g.dst, lines[i])
			}
		}
	} else {
		for i, match := range matchLine {
			if match {
				fmt.Fprintf(g.dst, "%d:%s\n", i+1, lines[i])
			} else if _, ok := printLines[i]; ok {
				fmt.Fprintf(g.dst, "%d-%s\n", i+1, lines[i])
			}
		}
	}
}

func GetContext(context, before, after int) (start, end int) {
	if context > 0 {
		start, end = context, context
	}

	if after > start {
		start = after
	}

	if before > end {
		end = before
	}
	return
}

func GetPattern(pattern string, fixed, ignoreCase bool) *regexp.Regexp {
	if fixed {
		pattern = regexp.QuoteMeta(pattern)
	}

	if ignoreCase {
		pattern = "(?i)" + pattern
	}

	return regexp.MustCompile(pattern)
}
