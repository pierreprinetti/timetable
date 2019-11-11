package interval_test

import (
	"testing"
	"time"

	"github.com/pierreprinetti/timetable/interval"
)

func TestInterval(t *testing.T) {
	mustParse := func(value string) time.Time {
		t, err := time.Parse(time.RFC3339, value)
		if err != nil {
			panic(err)
		}
		return t
	}

	t.Run("Interval with start and end dates and hours", func(t *testing.T) {
		var (
			// Outside the interval
			t0 = mustParse("2019-09-10T16:00:00Z")

			// Interval start
			t1 = mustParse("2019-10-10T14:00:00Z")

			// In interval
			t2 = mustParse("2019-11-05T15:12:33Z")

			// Interval end
			t3 = mustParse("2019-11-10T16:00:00Z")

			// Inside date interval, but outside hours
			t4 = mustParse("2019-11-05T05:12:33Z")
		)

		i := interval.New(
			interval.StartDate(t1),
			interval.EndDate(t3),
			interval.Clock(t1, t3),
		)

		if !i.Contains(t2) {
			t.Errorf("expected %v to contain %v", i, t2)
		}

		if i.Contains(t0) {
			t.Errorf("expected %v not to contain %v", i, t0)
		}

		if i.Contains(t4) {
			t.Errorf("expected %v not to contain %v", i, t4)
		}
	})

	t.Run("Interval with start and end dates, no hours", func(t *testing.T) {
		var (
			// Before the interval
			t0 = mustParse("2019-09-10T16:00:00Z")

			// Interval start
			t1 = mustParse("2019-10-10T14:00:00Z")

			// In interval
			t2 = mustParse("2019-11-05T15:12:33Z")

			// Interval end
			t3 = mustParse("2019-11-10T16:00:00Z")

			// Inside date interval, but outside hours
			t4 = mustParse("2019-11-05T05:12:33Z")

			// After the interval
			t5 = mustParse("2021-11-05T15:12:33Z")
		)

		i := interval.New(
			interval.StartDate(t1),
			interval.EndDate(t3),
		)

		if !i.Contains(t2) {
			t.Errorf("expected %v to contain %v", i, t2)
		}

		if i.Contains(t0) {
			t.Errorf("expected %v not to contain %v", i, t0)
		}

		// Expect all the interval to be valid if no hours are passed to New
		if !i.Contains(t4) {
			t.Errorf("expected %v to contain %v", i, t4)
		}

		if i.Contains(t5) {
			t.Errorf("expected %v not to contain %v", i, t5)
		}
	})

	t.Run("Interval with weekdays", func(t *testing.T) {
		i := interval.New(
			interval.Weekdays(time.Monday),
		)

		// Monday
		t1 := mustParse("2019-11-11T12:12:12Z")

		// Tuesday
		t2 := mustParse("2019-11-12T12:12:12Z")

		if !i.Contains(t1) {
			t.Errorf("expected %v to contain %v", i, t1)
		}

		if i.Contains(t2) {
			t.Errorf("expected %v not to contain %v", i, t2)
		}
	})
}
