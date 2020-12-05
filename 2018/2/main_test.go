package main

import "testing"

func Test_xplicates(t *testing.T) {
	tests := []struct {
		id    string
		dupl  int
		tripl int
	}{
		{"abcdef", 0, 0},
		{"bababc", 1, 1},
		{"abbcde", 1, 0},
		{"abcccd", 0, 1},
		{"aabcdd", 1, 0},
		{"abcdee", 1, 0},
		{"ababab", 0, 1},
	}
	for _, tt := range tests {
		t.Run(tt.id, func(t *testing.T) {

			got, got1 := xplicates(tt.id)

			if got != tt.dupl {
				t.Errorf("xplicates() got = %v, want %v", got, tt.dupl)
			}

			if got1 != tt.tripl {
				t.Errorf("xplicates() got1 = %v, want %v", got1, tt.tripl)
			}
		})
	}
}

func TestChecksum(t *testing.T) {
	ids := []string{
		"abcdef",
		"bababc",
		"abbcde",
		"abcccd",
		"aabcdd",
		"abcdee",
		"ababab",
	}

	if got, want := checksum(ids), 12; got != want {
		t.Errorf("Checksum() = %v want %v", got, want)
	}
}

func TestFindBoxes(t *testing.T) {
	ids := []string{
		"abcde",
		"fghij",
		"klmno",
		"pqrst",
		"fguij",
		"axcye",
		"wvxyz",
	}

	if got, want := findBoxes(ids), []string{"fghij", "fguij"}; !testEq(got, want) {
		t.Fatalf("findBoxes() = %+v, want %+v", got, want)
	}
}

func TestStringDiff(t *testing.T) {
	tests := []struct {
		id1  string
		id2  string
		diff int
	}{
		{"abcde", "abcde", 0},
		{"fghij", "fguij", 1},
		{"abcde", "axcye", 2},
	}
	for _, tt := range tests {
		t.Run(tt.id1+"-"+tt.id2, func(t *testing.T) {
			if got, want := stringDiff(tt.id1, tt.id2), tt.diff; got != want {
				t.Fatalf("stringDiff() = %v, want %v", got, want)
			}
		})
	}
}

func TestCommon(t *testing.T) {
	tests := []struct {
		id1    string
		id2    string
		common string
	}{
		{"abcde", "abcde", "abcde"},
		{"fghij", "fguij", "fgij"},
		{"abcde", "axcye", "ace"},
		{"abcde", "fghij", ""},
	}
	for _, tt := range tests {
		t.Run(tt.id1+"-"+tt.id2, func(t *testing.T) {
			if got, want := common(tt.id1, tt.id2), tt.common; got != want {
				t.Fatalf("common() = %v, want %v", got, want)
			}
		})
	}
}

func testEq(a, b []string) bool {

	// If one is nil, the other must also be nil.
	if (a == nil) != (b == nil) {
		return false
	}

	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
