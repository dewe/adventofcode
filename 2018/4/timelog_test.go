package main

import (
	"reflect"
	"testing"
	"time"
)

func TestNewTimelog(t *testing.T) {

	tests := []struct {
		input  string
		result timelog
	}{
		{
			"[1518-04-06 00:00] Guard #499 begins shift",
			timelog{time.Date(1518, 4, 6, 0, 0, 0, 0, time.UTC), "Guard #499 begins shift"},
		},
		{
			"[1518-04-06 00:09] falls asleep",
			timelog{time.Date(1518, 4, 6, 0, 9, 0, 0, time.UTC), "falls asleep"},
		},
		{
			"[1518-04-06 00:32] wakes up",
			timelog{time.Date(1518, 4, 6, 0, 32, 0, 0, time.UTC), "wakes up"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			if got, want := newTimelog(tt.input), tt.result; !reflect.DeepEqual(got, want) {
				t.Fatalf("newTimelog() = %+v, want %+v", got, want)
			}
		})
	}

}

func TestGuardID(t *testing.T) {

	tests := []struct {
		tl timelog
		id string
	}{
		{
			timelog{time.Date(1518, 4, 6, 0, 0, 0, 0, time.UTC), "Guard #499 begins shift"},
			"499",
		},
		{
			timelog{time.Date(1518, 4, 6, 0, 9, 0, 0, time.UTC), "falls asleep"},
			"",
		},
		{
			timelog{time.Date(1518, 4, 6, 0, 32, 0, 0, time.UTC), "wakes up"},
			"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.tl.event, func(t *testing.T) {
			if got, want := tt.tl.guardID(), tt.id; got != want {
				t.Fatalf("guardID() = %q, want %q", got, want)
			}
		})
	}

}
