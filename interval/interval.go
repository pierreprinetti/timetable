// package interval exposes the Interval type that contains information about
// "opening hours". An interval has an optional start date, an optional end
// date, optional start and end hours, and optionally excludes weekdays. For
// example:
//
//    t1, _ := time.Parse(time.RFC3339, "0000-00-00T10:00:00Z")
//    t2, _ := time.Parse(time.RFC3339, "0000-00-00T20:00:00Z")
//    openingHours := interval.New(
//        Clock(t1, t2),
//        Weekdays(
//            time.Monday,
//            time.Tuesday,
//            time.Wednesday,
//            time.Thursday,
//            time.Friday,
//    )
package interval

import (
	"fmt"
	"time"

	"github.com/pierreprinetti/timetable/clock"
	"github.com/pierreprinetti/timetable/week"
)

type Interval struct {
	StartDate time.Time
	EndDate   time.Time

	StartClock clock.Clock
	EndClock   clock.Clock

	Weekdays week.Week
}

func (i Interval) Contains(t time.Time) bool {
	if !i.Weekdays.Contains(t.Weekday()) {
		return false
	}

	if !i.StartDate.IsZero() && t.Before(i.StartDate) {
		return false
	}

	if !i.EndDate.IsZero() && t.After(i.EndDate) {
		return false
	}

	// Only consider hours if one of them is not zero
	if !i.StartClock.IsZero() || !i.StartClock.IsZero() {
		if c := clock.FromTime(t); c.Before(i.StartClock) || c.After(i.EndClock) {
			return false
		}
	}

	return true
}

func (i Interval) String() string {
	return fmt.Sprintf("From %s to %s, from %v to %v, on %v", i.StartDate, i.EndDate, i.StartClock, i.EndClock, i.Weekdays)
}

// New creates a new Interval, valid permanently by default. Options can be
// passed to restrict the validity of the interval.
//
// Examples.
//
// This interval is valid everyday from 2019-11-01 to 2019-11-07, from 7 to
// 9:30:
//
//    t1, _ := time.Parse(time.RFC3339, "2019-11-01T07:00:00Z")
//    t2, _ := time.Parse(time.RFC3339, "2019-11-07T09:30:00Z")
//    interval := New(
//        StartDate(t1),
//        EndDate(t2),
//        Clock(t1, t2),
//    )
//
//
// This interval is valid everyday from 2019-11-01 to 2019-11-06, from 7 to
// 9:30, except the first day, when will only start at 9:
//
//    t1, _ := time.Parse(time.RFC3339, "2019-11-01T09:00:00Z")
//    t2, _ := time.Parse(time.RFC3339, "2019-11-07T06:59:00Z")
//    c1, _ := time.Parse("15:04", "07:00")
//    c2, _ := time.Parse("15:04", "09:30")
//    interval := New(
//        StartDate(t1),
//        EndDate(t2),
//        Clock(c1, c2),
//    )
//
//
// This interval is valid every Monday and Wednesday of October, from 14 to 16:
//
//    t1, _ := time.Parse(time.RFC3339, "2019-10-01T14:00:00Z")
//    t2, _ := time.Parse(time.RFC3339, "2019-10-31T16:00:00Z")
//    interval := New(
//        StartDate(t1),
//        EndDate(t2),
//        Clock(t1, t2),
//        Weekdays(time.Monday, time.Wednesday),
//    )
//
//
// This interval is valid on Saturdays and Sundays in October:
//
//    t1, _ := time.Parse(time.RFC3339, "2019-10-01T00:00:00Z")
//    t2, _ := time.Parse(time.RFC3339, "2019-11-01T00:00:00Z")
//    interval := New(
//        StartDate(t1),
//        EndDate(t2),
//        Weekdays(time.Friday, time.Saturday, time.Sunday),
//    )
func New(options ...option) Interval {
	i := Interval{
		Weekdays: week.All,
	}

	for _, apply := range options {
		apply(&i)
	}

	return i
}
