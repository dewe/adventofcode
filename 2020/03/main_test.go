package main

import (
	"testing"
)

func Test_moduloHit(t *testing.T) {
	type args struct {
		row string
		x   int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"miss", args{"..##.......", 0}, 0},
		{"miss", args{"#...#...#..", 3}, 0},
		{"hit", args{".#....#..#.", 6}, 1},
		{"modulo miss", args{"..#.#...#.#", 12}, 0},
		{"modulo hit", args{".#...##..#.", 31}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := moduloHit(tt.args.row, tt.args.x); got != tt.want {
				t.Errorf("moduloHit() = %v, want %v", got, tt.want)
			}
		})
	}
}

var topology = []string{
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#",
}

func Test_run(t *testing.T) {
	type args struct {
		slope []string
		right int
		down  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"1r1d", args{topology, 1, 1}, 2},
		{"3r1d", args{topology, 3, 1}, 7},
		{"5r1d", args{topology, 5, 1}, 3},
		{"7r1d", args{topology, 7, 1}, 4},
		{"1r2d", args{topology, 1, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := run(tt.args.slope, tt.args.right, tt.args.down); got != tt.want {
				t.Errorf("run() = %v, want %v", got, tt.want)
			}
		})
	}
}
