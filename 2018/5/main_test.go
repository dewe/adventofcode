package main

import "testing"

func TestReact(t *testing.T) {

	tests := []struct {
		in  string
		out string
	}{
		{in: "aA", out: ""},
		{in: "abBA", out: ""},
		{in: "ab", out: "ab"},
		{in: "dabAcCaCBAcCcaDA", out: "dabCBAcaDA"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {

			units, _ := react(tt.in)

			if got, want := units, tt.out; got != want {
				t.Fatalf("reduce() = %v, want %v", got, want)
			}

		})
	}

}

func TestUniqueChars(t *testing.T) {

	tests := []struct {
		in  string
		out string
	}{
		{in: "aA", out: "a"},
		{in: "abBA", out: "ab"},
		{in: "dabAcCaCBAcCcaDA", out: "dabc"},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {

			if got, want := uniqueChars(tt.in), tt.out; got != want {
				t.Fatalf("reduce() = %v, want %v", got, want)
			}

		})
	}

}
