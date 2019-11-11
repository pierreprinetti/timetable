package week_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/pierreprinetti/timetable/week"
)

func TestWeekContains(t *testing.T) {
	for _, tc := range [...]struct {
		week     []time.Weekday
		contains []time.Weekday
	}{
		{
			week:     []time.Weekday{time.Sunday},
			contains: []time.Weekday{time.Sunday},
		},
		{
			week: []time.Weekday{
				time.Monday,
				time.Tuesday,
				time.Wednesday,
			},
			contains: []time.Weekday{time.Tuesday},
		},
	} {
		t.Run(fmt.Sprintf("Week %q contains %q", tc.week, tc.contains), func(t *testing.T) {
			if !week.New(tc.week...).Contains(tc.contains...) {
				t.Errorf("expected %q to contain %q", tc.week, tc.contains)
			}
		})
	}

	for _, tc := range [...]struct {
		week        []time.Weekday
		containsNot []time.Weekday
	}{
		{
			week:        []time.Weekday{},
			containsNot: []time.Weekday{time.Sunday},
		},
		{
			week:        []time.Weekday{time.Monday},
			containsNot: []time.Weekday{time.Sunday},
		},
		{
			week: []time.Weekday{time.Wednesday},
			containsNot: []time.Weekday{
				time.Wednesday,
				time.Saturday},
		},
	} {
		t.Run(fmt.Sprintf("Week %q does not contain %q", tc.week, tc.containsNot), func(t *testing.T) {
			if week.New(tc.week...).Contains(tc.containsNot...) {
				t.Errorf("expected %q not to contain %q", tc.week, tc.containsNot)
			}
		})
	}
}

func TestWeekAll(t *testing.T) {
	t.Run("All contains all the weekdays", func(t *testing.T) {
		if !week.All.Contains(
			time.Monday,
			time.Tuesday,
			time.Wednesday,
			time.Thursday,
			time.Friday,
			time.Saturday,
			time.Sunday,
		) {
			t.Error("found false")
		}
	})
}
