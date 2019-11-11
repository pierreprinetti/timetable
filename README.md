# timetable

A Go type for recurring events, like opening hours.

## Example

```Go
package main

import (
	"fmt"
	"time"

	"github.com/pierreprinetti/timetable"
	"github.com/pierreprinetti/timetable/interval"
)

func main() {
	// Every chunk of recurring schedule is an interval. Here we create the
	// interval for the morning opening hours.
	t1, _ := time.Parse(time.RFC3339, "0001-01-01T08:00:00Z")
	t2, _ := time.Parse(time.RFC3339, "0001-01-01T12:00:00Z")
	amOpeningHours := interval.New(
		interval.Clock(t1, t2),
		interval.Weekdays(
			time.Monday,
			time.Tuesday,
			time.Wednesday,
			time.Thursday,
			time.Friday,
			time.Saturday,
		),
	)

	// Here we craete the interval for the afternoon opening hours.
	t3, _ := time.Parse(time.RFC3339, "0001-01-01T14:30:00Z")
	t4, _ := time.Parse(time.RFC3339, "0001-01-01T19:30:00Z")
	pmOpeningHours := interval.New(
		interval.Clock(t3, t4),
		interval.Weekdays(
			time.Monday,
			time.Tuesday,
			time.Wednesday,
			time.Thursday,
			time.Friday,
			time.Saturday,
		),
	)

	// And now, the timetable: a slice of Intervals.
	openingHours := timetable.Timetable{
		amOpeningHours,
		pmOpeningHours,
	}

	t0, _ := time.Parse(time.RFC3339, "2019-11-08T09:00:00Z")
	if openingHours.Contains(t0) {
		fmt.Println("True!")
	}
}
```
