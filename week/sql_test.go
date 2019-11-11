package week_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/pierreprinetti/timetable/week"
)

func TestWeekScan(t *testing.T) {
	for _, tc := range [...]struct {
		input interface{}
		want  week.Week
	}{
		{
			"1",
			week.New(time.Sunday),
		},
		{
			"10",
			week.New(time.Monday),
		},
		{
			"11",
			week.New(time.Monday, time.Sunday),
		},
		{
			"1111111",
			week.All,
		},
	} {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			var have week.Week

			if err := have.Scan(tc.input); err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if have != tc.want {
				t.Errorf("expected %v to equal %v", have, tc.want)
			}
		})
	}

	for _, tc := range [...]interface{}{
		"2",
		"ciao",
		"",
		" ",
		11,
	} {
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			var have week.Week

			if err := have.Scan(tc); err == nil {
				t.Errorf("expected error, got nil")
			}
		})
	}
}

func TestWeekValue(t *testing.T) {
	for _, tc := range [...]struct {
		w week.Week
		v string
	}{
		{
			week.New(time.Sunday),
			"1",
		},
		{
			week.New(time.Saturday, time.Sunday),
			"1000001",
		},
		{
			week.New(time.Wednesday, time.Saturday, time.Sunday),
			"1001001",
		},
		{
			week.New(),
			"0",
		},
		{
			week.All,
			"1111111",
		},
	} {
		t.Run(tc.v, func(t *testing.T) {
			have, err := tc.w.Value()
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if tc.v != have {
				t.Errorf("expected %v, found %v", tc.v, have)
			}
		})
	}
}

func TestWeekString(t *testing.T) {
	for _, tc := range [...]struct {
		w week.Week
		v string
	}{
		{
			week.New(time.Sunday),
			"Sundays",
		},
		{
			week.New(time.Saturday, time.Sunday),
			"Sundays, Saturdays",
		},
		{
			week.New(time.Wednesday, time.Saturday, time.Sunday),
			"Sundays, Wednesdays, Saturdays",
		},
		{
			week.New(),
			"",
		},
		{
			week.All,
			"Sundays, Mondays, Tuesdays, Wednesdays, Thursdays, Fridays, Saturdays",
		},
	} {
		t.Run(tc.v, func(t *testing.T) {
			have := tc.w.String()

			if tc.v != have {
				t.Errorf("expected %v, found %v", tc.v, have)
			}
		})
	}
}
