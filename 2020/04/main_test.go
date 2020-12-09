package main

import (
	"io"
	"reflect"
	"strings"
	"testing"
)

// // valid
// ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
// byr:1937 iyr:2017 cid:147 hgt:183cm

// // invalid
// iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
// hcl:#cfa07d byr:1929

// // valid
// hcl:#ae17e1 iyr:2013
// eyr:2024
// ecl:brn pid:760753108 byr:1931
// hgt:179cm

// invalid
// hcl:#cfa07d eyr:2025 pid:166559648
// iyr:2011 ecl:brn hgt:59in

var passports = []string{
	"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm",
	"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929",
	"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm",
	"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in",
}

func Test_countPassports(t *testing.T) {
	type args struct {
		passports []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"example", args{passports}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countPassports(tt.args.passports); got != tt.want {
				t.Errorf("countPassports() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_scanBlankLines(t *testing.T) {
	type args struct {
		file io.Reader
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{"blanks line", args{strings.NewReader("a\n\nb\n\nc")}, []string{"a", "b", "c"}},
		{"linebreak", args{strings.NewReader("a\nb")}, []string{"a b"}},
		{"ending linebreak", args{strings.NewReader("a\n")}, []string{"a"}},
		// {"double blanks", args{strings.NewReader("a\n\n\nb")}, []string{"a","b"}},
		{"full monty", args{strings.NewReader("a\nb\n\nc\n\nd\ne\nf")}, []string{"a b", "c", "d e f"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := scanBlankLines(tt.args.file); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scanBlankLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isValid(t *testing.T) {
	type args struct {
		passport string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// {"valid keys", args{"ecl:gry pid:860033327 eyr:2020 hcl:#fffffd byr:1937 iyr:2017 cid:147 hgt:183cm"}, 1},
		// {"invalid - missing hgt", args{"iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884 hcl:#cfa07d byr:1929"}, 0},
		// {"optional key", args{"hcl:#ae17e1 iyr:2013 eyr:2024 ecl:brn pid:760753108 byr:1931 hgt:179cm"}, 1},
		// {"invalid - missing key byd", args{"hcl:#cfa07d eyr:2025 pid:166559648 iyr:2011 ecl:brn hgt:59in"}, 0},

		{"valid values", args{"pid:087499704 hgt:74in ecl:grn iyr:2012 eyr:2030 byr:1980 hcl:#623a2f"}, 1},
		{"valid values", args{"eyr:2029 ecl:blu cid:129 byr:1989 iyr:2014 pid:896056539 hcl:#a97842 hgt:165cm"}, 1},
		{"valid values", args{"hcl:#888785 hgt:164cm byr:2001 iyr:2015 cid:88 pid:545766238 ecl:hzl eyr:2022"}, 1},
		{"valid values", args{"iyr:2010 hgt:158cm hcl:#b6652a ecl:blu byr:1944 eyr:2021 pid:093154719"}, 1},

		{"valid values", args{"eyr:1972 cid:100 hcl:#18171d ecl:amb hgt:170 pid:186cm iyr:2018 byr:1926"}, 0},
		{"valid values", args{"iyr:2019 hcl:#602927 eyr:1967 hgt:170cm ecl:grn pid:012533040 byr:1946"}, 0},
		{"valid values", args{"hcl:dab227 iyr:2012 ecl:brn hgt:182cm pid:021572410 eyr:2020 byr:1992 cid:277"}, 0},
		{"valid values", args{"hgt:59cm ecl:zzz eyr:2038 hcl:74454a iyr:2023 pid:3556412378 byr:2007"}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isValid(tt.args.passport); got != tt.want {
				t.Errorf("isValid() = %v, want %v", got, tt.want)
			}
		})
	}
}
