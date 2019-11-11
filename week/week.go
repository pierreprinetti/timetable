package week

import (
	"strings"
	"time"
)

const All = Week(127)

type Week byte

func New(weekdays ...time.Weekday) Week {
	var w Week
	for _, weekday := range weekdays {
		w = w | 1<<uint(weekday)
	}
	return w
}

func (w Week) String() string {
	var weekdays []string
	for i := range [7]struct{}{} {
		weekday := time.Weekday(i)
		if w.Contains(weekday) {
			weekdays = append(weekdays, weekday.String()+"s")
		}
	}

	return strings.Join(weekdays, ", ")
}

func (w Week) Contains(weekdays ...time.Weekday) bool {
	return w == w|New(weekdays...)
}
