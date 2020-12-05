package main

type guard struct {
	id       int
	timelogs []timelog
	sleep    [60]int
}

func (g guard) addTimelog(tl timelog) guard {

	if tl.event == "wakes up" {
		start := last(g.timelogs).minute()
		end := tl.minute()

		for i := start; i < end; i++ {
			g.sleep[i]++
		}
	}

	g.timelogs = append(g.timelogs, tl)

	return g
}

func last(tls []timelog) timelog {
	return tls[len(tls)-1]
}

func (g guard) maxAsleep() int {
	var max int
	for _, val := range g.sleep {
		if val > max {
			max = val
		}
	}
	return max
}
