package main

import "testing"

func TestIsValid(t *testing.T) {
	type args struct {
		password string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"valid", args{"1-3 a: abcde"}, true},
		{"invalid, too few", args{"1-3 b: cdefg"}, false},
		{"invalid, too many", args{"2-9 c: ccccccccc"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsValid(tt.args.password); got != tt.want {
				t.Errorf("IsValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
