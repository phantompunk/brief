package command

import (
	"fmt"
	"time"
)

type weekYear struct {
	Week int
	Year int
	Mon  string
	Tue  string
	Wed  string
	Thu  string
	Fri  string
}

var days = map[int]int{
	0: -1,
	1: 0,
	2: 1,
	3: 2,
	4: 3,
	5: 4,
	6: -2,
}

func getDates(start time.Time) *weekYear {
	year, week := start.ISOWeek()

	firstDayOfWeek := start.AddDate(0, 0, -days[int(start.Weekday())])
	_, m, d := firstDayOfWeek.Date()
	monday := fmt.Sprintf("%d.%d", m, d)

	_, m, d = firstDayOfWeek.AddDate(0, 0, 1).Date()
	tuesday := fmt.Sprintf("%d.%d", m, d)

	_, m, d = firstDayOfWeek.AddDate(0, 0, 2).Date()
	wednesday := fmt.Sprintf("%d.%d", m, d)

	_, m, d = firstDayOfWeek.AddDate(0, 0, 3).Date()
	thursday := fmt.Sprintf("%d.%d", m, d)

	_, m, d = firstDayOfWeek.AddDate(0, 0, 4).Date()
	friday := fmt.Sprintf("%d.%d", m, d)

	return &weekYear{week, year, monday, tuesday, wednesday, thursday, friday}
}
