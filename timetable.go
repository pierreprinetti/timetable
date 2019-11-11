// package timetable exposes the Timetable type that contains information about
// a recurring schedule. A Timetable is the union of zero or more
// interval.Interval. For example:
//
//	t1, _ := time.Parse(time.RFC3339, "0001-01-01T08:00:00Z")
//	t2, _ := time.Parse(time.RFC3339, "0001-01-01T12:00:00Z")
//	amOpeningHours := interval.New(
//		interval.Clock(t1, t2),
//		interval.Weekdays(
//			time.Monday,
//			time.Tuesday,
//			time.Wednesday,
//			time.Thursday,
//			time.Friday,
//			time.Saturday,
//		),
//	)
//
//	t3, _ := time.Parse(time.RFC3339, "0001-01-01T14:30:00Z")
//	t4, _ := time.Parse(time.RFC3339, "0001-01-01T19:30:00Z")
//	pmOpeningHours := interval.New(
//		interval.Clock(t3, t4),
//		interval.Weekdays(
//			time.Monday,
//			time.Tuesday,
//			time.Wednesday,
//			time.Thursday,
//			time.Friday,
//			time.Saturday,
//		),
//	)
//
//	openingHours := timetable.Timetable{
//		amOpeningHours,
//		pmOpeningHours,
//	}
package timetable

import (
	"time"

	"github.com/pierreprinetti/timetable/interval"
)

// Timetable defines a schedule by observing several intervals.
type Timetable []interval.Interval

// Contains returns true if at least one of the intervals contains the instant
// t.
func (tt Timetable) Contains(t time.Time) bool {
	for _, interv := range tt {
		if interv.Contains(t) {
			return true
		}
	}
	return false
}
