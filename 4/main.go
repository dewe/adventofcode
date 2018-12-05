package main

import (
	"bufio"
	"os"
	"sort"
	"strconv"
)

// [1518-04-06 00:00] Guard #499 begins shift
// [1518-04-06 00:09] falls asleep
// [1518-04-06 00:32] wakes up
// [1518-04-06 00:43] falls asleep
// [1518-04-06 00:55] wakes up

func main() {
	logs := readTimelogs("input.txt")
	guardlogs := logsByGuard(logs)
	sleeplogs := sleepByGuard(guardlogs)

	// brutto sleeper
	var sleeper int
	var maxSleep int
	for id, sl := range sleeplogs {
		if sumSleep(sl) > maxSleep {
			maxSleep = sumSleep(sl)
			sleeper = id
		}
	}
	minute, _ := maxMinute(sleeplogs[sleeper])

	println("Guard:", sleeper)
	println("Minute:", minute)
	println("A1:", sleeper*minute)

	// maxMinuteSleeper
	sleeper = 0
	maxSleep = 0
	minute = 0
	for id, sl := range sleeplogs {
		if m, max := maxMinute(sl); max > maxSleep {
			sleeper, maxSleep, minute = id, max, m
		}
	}

	println("Guard:", sleeper)
	println("Minute:", minute)
	println("A2:", sleeper*minute)

}

func maxMinute(sl sleeplog) (int, int) {
	var minute int
	var max int

	for m, val := range sl {
		if val > max {
			max, minute = val, m
		}
	}

	return minute, max
}

func sumSleep(sl sleeplog) int {
	var sum int

	for _, v := range sl {
		sum += v
	}

	return sum
}

type sleeplog [60]int

func sleepByGuard(guardlogs map[int]timelogs) map[int]sleeplog {
	var sleeplogs = map[int]sleeplog{}

	for id, logs := range guardlogs {
		sleeplogs[id] = sleepFromLogs(logs)
	}

	return sleeplogs
}

func sleepFromLogs(logs timelogs) [60]int {
	var prev timelog
	var curr timelog
	var sleep [60]int

	for _, curr = range logs {

		if curr.event == "wakes up" {
			start := prev.minute()
			end := curr.minute()

			for i := start; i < end; i++ {
				sleep[i]++
			}
		}

		prev = curr
	}

	return sleep
}

func logsByGuard(logs []timelog) map[int]timelogs {
	var m = map[int]timelogs{}
	var gid int

	for _, tl := range logs {
		// [1518-04-06 00:00] Guard #499 begins shift
		if s := tl.guardID(); s != "" {
			gid, _ = strconv.Atoi(s)
		}

		m[gid] = append(m[gid], tl)
	}

	return m
}

func byGuard(timelogs []timelog) map[int]guard {
	var m = map[int]guard{}
	var gid int

	for _, tl := range timelogs {
		// [1518-04-06 00:00] Guard #499 begins shift
		if s := tl.guardID(); s != "" {
			id, _ := strconv.Atoi(s)
			gid = id
		}

		guard := m[gid]
		guard.id = gid
		m[gid] = guard.addTimelog(tl)
	}

	return m
}

func readTimelogs(name string) []timelog {
	timelogs := []timelog{}
	for _, l := range readSorted(name) {
		timelogs = append(timelogs, newTimelog(l))
	}
	return timelogs
}

func readSorted(name string) []string {
	lines := readLines(name)

	sort.Strings(lines)

	return lines
}

func readLines(name string) []string {
	lines := []string{}

	file, _ := os.Open(name)
	s := bufio.NewScanner(file)
	for s.Scan() {
		lines = append(lines, s.Text())
	}
	return lines
}
