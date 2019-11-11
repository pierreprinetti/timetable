package interval

import (
	"time"

	"github.com/pierreprinetti/timetable/clock"
	"github.com/pierreprinetti/timetable/week"
)

type option func(*Interval)

// StartDate can be passed as an option to New to start the interval at a specific date and time.
//
// Example:
//
//    t1, _ := time.Parse(time.RFC3339, "2019-11-05T07:00:00Z")
//    interval := New(StartDate(t1))
func StartDate(t time.Time) option {
	return func(i *Interval) {
		i.StartDate = t
	}
}

// EndDate can be passed as an option to New to end the interval at a specific date and time.
//
// Example:
//
//    t2, _ := time.Parse(time.RFC3339, "2019-11-07T09:30:00Z")
//    interval := New(EndDate(t2))
func EndDate(t time.Time) option {
	return func(i *Interval) {
		i.EndDate = t
	}
}

// Clock can be passed as an option to New to restrict the interval to specific hours.
//
// Example:
//
//    t1, _ := time.Parse("15:04", "14:00")
//    t2, _ := time.Parse("15:04", "16:00")
//    interval := New(Clock(t1, t2))
func Clock(start, end time.Time) option {
	return func(i *Interval) {
		i.StartClock = clock.FromTime(start)
		i.EndClock = clock.FromTime(end)
	}
}

// Weekdays can be passed as an option to New to restrict the interval to specific weekdays.
//
// Example:
//
//    interval := New(Weekdays(time.Monday, time.Wednesday))
func Weekdays(weekdays ...time.Weekday) option {
	return func(i *Interval) {
		i.Weekdays = week.New(weekdays...)
	}
}
