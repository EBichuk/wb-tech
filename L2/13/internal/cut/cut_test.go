package cut

import (
	"reflect"
	"testing"
)

func TestParseFields(t *testing.T) {
	type args struct {
		fields string
	}
	tests := []struct {
		name    string
		args    args
		want    []int
		wantErr bool
	}{
		{
			name:    "success: numbers are separated by commas",
			args:    args{fields: "1,3"},
			want:    []int{1, 3},
			wantErr: false,
		},
		{
			name:    "success: numbers are specified by a range",
			args:    args{fields: "1-3"},
			want:    []int{1, 2, 3},
			wantErr: false,
		},
		{
			name:    "success: numbers are specified separated by commas and in a range",
			args:    args{fields: "1,2-4"},
			want:    []int{1, 2, 3, 4},
			wantErr: false,
		},
		{
			name:    "borderline: numbers are in the range (start == end)",
			args:    args{fields: "2-2"},
			want:    []int{2},
			wantErr: false,
		},
		{
			name:    "borderline: numbers are in the wrong range (end < strart)",
			args:    args{fields: "4-2"},
			want:    nil,
			wantErr: false,
		},
		{
			name:    "error: with zero",
			args:    args{fields: "0,2"},
			want:    nil,
			wantErr: true,
		},
		{
			name:    "error: with letters",
			args:    args{fields: "1-o"},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseFields(tt.args.fields)

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFields() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFields() = %v, want %v", got, tt.want)
			}
		})
	}
}
