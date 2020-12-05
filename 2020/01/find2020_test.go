package main

import "testing"

func TestFindTwo2020(t *testing.T) {
	type args struct {
		expenses []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"valid",
			args{[]int{1721, 979, 366, 299, 675, 1456}},
			514579,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindTwo2020(tt.args.expenses); got != tt.want {
				t.Errorf("Find2020() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFindThree2020(t *testing.T) {
	type args struct {
		expenses []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"valid",
			args{[]int{1721, 979, 366, 299, 675, 1456}},
			241861950,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FindThree2020(tt.args.expenses); got != tt.want {
				t.Errorf("Find2020() = %v, want %v", got, tt.want)
			}
		})
	}
}
