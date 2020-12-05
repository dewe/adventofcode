package main

import (
	"strings"
	"time"
)

type timelog struct {
	t     time.Time
	event string
}

type timelogs []timelog

// [1518-04-06 00:09] falls asleep
func newTimelog(s string) timelog {
	t, _ := time.Parse("2006-01-02 15:04", s[1:17])
	return timelog{t, s[19:]}
}

func (tl timelog) guardID() string {
	parts := strings.Split(tl.event, " ")
	if parts[0] == "Guard" {
		return parts[1][1:]
	}
	return ""
}

func (tl timelog) minute() int {
	_, minute, _ := tl.t.Clock()
	return minute
}
