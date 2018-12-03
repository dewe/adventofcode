package main

import (
	"reflect"
	"testing"
)

func Test_newClaim(t *testing.T) {
	tests := []struct {
		s    string
		want claim
	}{
		{"#1 @ 1,3: 4x4", claim{"1", 1, 3, 4, 4}},
		{"#259 @ 558,939: 26x10", claim{"259", 558, 939, 26, 10}},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			if got := newClaim(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("newClaim() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTiles(t *testing.T) {
	tests := []struct {
		desc  string
		claim claim
		tiles []tile
	}{
		{"#1 @ 1,3: 1x1", claim{"1", 1, 3, 1, 1}, []tile{{1, 3}}},
		{"#1 @ 9,29: 2x2", claim{"1", 9, 29, 2, 2}, []tile{{9, 29}, {9, 30}, {10, 29}, {10, 30}}},
	}
	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {

			if got, want := tt.claim.tiles(), tt.tiles; !reflect.DeepEqual(got, want) {
				t.Errorf("tiles() = %v, want %v", got, want)
			}
		})
	}
}
