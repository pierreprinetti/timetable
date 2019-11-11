package timetable_test

import (
	"testing"
	"time"

	"github.com/pierreprinetti/timetable"
	"github.com/pierreprinetti/timetable/interval"
)

func TestTimetableExample(t *testing.T) {
	mustParse := func(t time.Time, err error) time.Time {
		if err != nil {
			panic(err)
		}
		return t
	}

	t1 := mustParse(time.Parse(time.RFC3339, "0001-01-01T08:00:00Z"))
	t2 := mustParse(time.Parse(time.RFC3339, "0001-01-01T12:00:00Z"))
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

	t3 := mustParse(time.Parse(time.RFC3339, "0001-01-01T14:30:00Z"))
	t4 := mustParse(time.Parse(time.RFC3339, "0001-01-01T19:30:00Z"))
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

	openingHours := timetable.Timetable{
		amOpeningHours,
		pmOpeningHours,
	}

	for _, tc := range [...]time.Time{
		t1,
		t2,
		t3,
		t4,
		mustParse(time.Parse(time.RFC3339, "0001-01-01T08:00:01Z")),
		mustParse(time.Parse(time.RFC3339, "2019-11-08T08:00:00Z")),
		mustParse(time.Parse(time.RFC3339, "2019-11-08T09:00:00Z")),
		mustParse(time.Parse(time.RFC3339, "1997-01-31T15:15:15Z")),
	} {
		if !openingHours.Contains(tc) {
			t.Errorf("expected to contain %s", tc)
		}
	}

	for _, tc := range [...]time.Time{
		mustParse(time.Parse(time.RFC3339, "1997-01-31T07:59:15Z")),
	} {
		if openingHours.Contains(tc) {
			t.Errorf("expected not to contain %s", tc)
		}
	}
}
