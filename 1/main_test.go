package main

import "testing"

func TestFindFirstRepeat(t *testing.T) {

	if got, want := findFirstRepeat([]int{1, -1}), 0; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	if got, want := findFirstRepeat([]int{3, 3, 4, -2, -4}), 10; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	if got, want := findFirstRepeat([]int{-6, 3, 8, 5, -6}), 5; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

	if got, want := findFirstRepeat([]int{7, 7, -2, -7, -4}), 14; got != want {
		t.Fatalf("got %v, want %v", got, want)
	}

}
