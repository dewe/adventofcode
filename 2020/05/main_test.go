package main

import "testing"

func Test_seatId(t *testing.T) {
	type args struct {
		code string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{ "row 44, column 5, seat ID 357.", args{"FBFBBFFRLR"}, 357}, 
		{ "row 70, column 7, seat ID 567.", args{"BFFFBBFRRR"}, 567}, 
		{ "row 14, column 7, seat ID 119.", args{"FFFBBBFRRR"}, 119}, 
		{ "row 102, column 4, seat ID 820.", args{"BBFFBBFRLL"}, 820}, 
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := seatId(tt.args.code); got != tt.want {
				t.Errorf("seatId() = %v, want %v", got, tt.want)
			}
		})
	}
}
