package main

import (
	"reflect"
	"testing"
)

func TestUnpackingString(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name:    "success: all letter unpacked",
			args:    args{str: "a1b4v5"},
			want:    "abbbbvvvvv",
			wantErr: false,
		},
		{
			name:    "success: big numbers",
			args:    args{str: "d12v13"},
			want:    "ddddddddddddvvvvvvvvvvvvv",
			wantErr: false,
		},
		{
			name:    "success: with 1",
			args:    args{str: "a1b1c1d1e1"},
			want:    "abcde",
			wantErr: false,
		},
		{
			name:    "success: without number in the end",
			args:    args{str: "a2b2c2d2e"},
			want:    "aabbccdde",
			wantErr: false,
		},
		{
			name:    "borderline: with 0",
			args:    args{str: "a0b0c0d0e0"},
			want:    "abcde",
			wantErr: false,
		},
		{
			name:    "borderline: with 0",
			args:    args{str: "a03bcde"},
			want:    "aaabcde",
			wantErr: false,
		},
		{
			name:    "borderline: without numbers",
			args:    args{str: "abcde"},
			want:    "abcde",
			wantErr: false,
		},
		{
			name:    "borderline: empty string",
			args:    args{str: ""},
			want:    "",
			wantErr: false,
		},
		{
			name:    "borderline: with minus",
			args:    args{str: "f-5s"},
			want:    "f-----s",
			wantErr: false,
		},
		{
			name:    "error: only numbers",
			args:    args{str: "45"},
			want:    "",
			wantErr: true,
		},
		{
			name:    "success: with escape sub",
			args:    args{str: "qwe\\4\\5"},
			want:    "qwe45",
			wantErr: false,
		},
		{
			name:    "success: with escape 0",
			args:    args{str: "qwe\\05"},
			want:    "qwe00000",
			wantErr: false,
		},
		{
			name:    "error: escape in the end",
			args:    args{str: "qwex5\\"},
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := unpackingString(tt.args.str)

			if (err != nil) != tt.wantErr {
				t.Errorf("unpackingString() error = %v, wantErr %v", err, tt.wantErr)

				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unpackingString() = %v, want %v", got, tt.want)
			}
		})
	}
}
