package main

import (
	"reflect"
	"testing"
)

func Test_parsePolicy(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name string
		args args
		want policy
	}{
		{"single digit limits", args{"2-5 a"}, policy{2, 5, "a"}},
		{"double digit limits", args{"12-55 b"}, policy{12, 55, "b"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parsePolicy(tt.args.t); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parsePolicy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_pw_isValid(t *testing.T) {
	type fields struct {
		policy   policy
		password string
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
	}{
		{ "valid", fields{ policy{2,5,"a"}, "aaa"}, true},
		{ "invalid", fields{policy{2,5, "a"}, "a"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &pw{
				policy:   tt.fields.policy,
				password: tt.fields.password,
			}
			if got := p.isValid(); got != tt.want {
				t.Errorf("pw.isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}

