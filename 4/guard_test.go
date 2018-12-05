package main

import (
	"reflect"
	"testing"
	"time"
)

func TestMaxAsleep(t *testing.T) {
	g := guard{sleep: [60]int{3: 1, 4: 3, 15: 2}}

	if got, want := g.maxAsleep(), 3; got != want {
		t.Fatalf("g.maxAsleep() = %v, want %v", got, want)
	}
}

func TestAddTimelog(t *testing.T) {

	t.Run("falls asleep", func(t *testing.T) {
		g := guard{
			timelogs: []timelog{timelog{moment("00:00"), "Guard #1 begins shift"}},
		}

		g2 := g.addTimelog(timelog{moment("00:10"), "falls asleep"})

		if got, want := len(g2.timelogs), 2; got != want {
			t.Fatalf("len() = %v, want %v", got, want)
		}

		if got, want := g.sleep, [60]int{}; !reflect.DeepEqual(got, want) {
			t.Fatalf("g.sleep() = %v, want %v", got, want)
		}
	})

	t.Run("wakes up", func(t *testing.T) {
		g1 := guard{
			timelogs: []timelog{
				timelog{moment("00:00"), "Guard #1 begins shift"},
				timelog{moment("00:10"), "falls asleep"},
			},
		}

		g2 := g1.addTimelog(timelog{moment("00:12"), "wakes up"})

		if got, want := len(g2.timelogs), 3; got != want {
			t.Fatalf("len() = %v, want %v", got, want)
		}

		if got, want := g2.sleep, [60]int{10: 1, 11: 1}; !reflect.DeepEqual(got, want) {
			t.Fatalf("g.sleep() = %v, want %v", got, want)
		}
	})

}

func moment(s string) time.Time {
	t, _ := time.Parse("15:04", s)
	return t
}
