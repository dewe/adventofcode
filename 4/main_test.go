package main

import (
	"reflect"
	"testing"
	"time"
)

func TestSleepLog(t *testing.T) {
	logs := timelogs{
		timelog{moment("00:00"), "Guard #10 begins shift"},
		timelog{moment("00:05"), "falls asleep"},
		timelog{moment("00:25"), "wakes up"},
		timelog{moment("00:30"), "falls asleep"},
		timelog{moment("00:55"), "wakes up"},
		timelog{moment("00:00"), "Guard #10 begins shift"},
		timelog{moment("00:24"), "falls asleep"},
		timelog{moment("00:29"), "wakes up"},
	}

	sleep := sleepLog(logs)
	want := [60]int{
		0, 0, 0, 0, 0, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 2, 1, 1, 1, 1, 0,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 1, 1, 1, 1, 1,
		1, 1, 1, 1, 1, 0, 0, 0, 0, 0,
	}

	if !reflect.DeepEqual(sleep, want) {
		t.Fatalf("sleep = %v, want %v", sleep, want)
	}
}

func TestLogsByGuard(t *testing.T) {
	logs := []timelog{
		timelog{time.Date(1518, 4, 6, 0, 0, 0, 0, time.UTC), "Guard #499 begins shift"},
		timelog{time.Date(1518, 4, 6, 0, 9, 0, 0, time.UTC), "falls asleep"},
		timelog{time.Date(1518, 4, 6, 0, 32, 0, 0, time.UTC), "wakes up"},
		timelog{time.Date(1518, 4, 7, 0, 1, 0, 0, time.UTC), "Guard #50 begins shift"},
		timelog{time.Date(1518, 4, 7, 0, 19, 0, 0, time.UTC), "falls asleep"},
	}

	guardlogs := logsByGuard(logs)

	if got, want := len(guardlogs), 2; got != want {
		t.Fatalf("len() = %v, want %v", got, want)
	}

	if got, want := len(guardlogs[499]), 3; got != want {
		t.Fatalf("len(499) = %v, want %v", got, want)
	}
}

func TestByGuard(t *testing.T) {
	timelogs := []timelog{
		timelog{time.Date(1518, 4, 6, 0, 0, 0, 0, time.UTC), "Guard #499 begins shift"},
		timelog{time.Date(1518, 4, 6, 0, 9, 0, 0, time.UTC), "falls asleep"},
		timelog{time.Date(1518, 4, 6, 0, 32, 0, 0, time.UTC), "wakes up"},
		timelog{time.Date(1518, 4, 6, 0, 33, 0, 0, time.UTC), "Guard #50 begins shift"},
	}

	if got, want := len(byGuard(timelogs)), 2; got != want {
		t.Fatalf("len() = %v, want %v", got, want)
	}

}
