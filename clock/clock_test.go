package clock_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/pierreprinetti/timetable/clock"
)

func TestNewUTC(t *testing.T) {
	if c1, c2 := clock.NewUTC(7, 59), clock.NewUTC(8, 0); !c1.Before(c2) {
		t.Errorf("expected %v to be before %v", c1, c2)
	}

	if c1, c2 := clock.NewUTC(8, 1), clock.NewUTC(8, 0); !c1.After(c2) {
		t.Errorf("expected %v to be after %v", c1, c2)
	}
}

func TestIsZero(t *testing.T) {
	var c0 clock.Clock
	if !c0.IsZero() {
		t.Errorf("expected %v to be zero", c0)
	}

	if c2 := clock.NewUTC(0, 1); c2.IsZero() {
		t.Errorf("expected %v not to be zero", c2)
	}
}

func TestBeforeAfter(t *testing.T) {
	mustParse := func(value string) time.Time {
		t, err := time.Parse(time.RFC3339, value)
		if err != nil {
			panic(err)
		}
		return t
	}

	for i, tc := range [...]struct {
		t1, t2                string
		equals, before, after bool
	}{
		{
			"2019-11-08T09:00:00Z", "2019-11-08T09:00:00Z",
			true, false, false,
		},
		{
			"2019-11-08T09:00:00Z", "1991-09-14T09:00:00Z",
			true, false, false,
		},
		{
			"2019-11-08T08:00:00Z", "2019-11-08T09:00:00Z",
			false, true, false,
		},
		{
			"2019-11-08T08:00:00Z", "2019-11-08T08:01:00Z",
			false, true, false,
		},
		{
			"2019-11-08T08:00:00Z", "2019-11-08T00:01:00Z",
			false, false, true,
		},
		{
			"2019-11-08T08:00:00Z", "2019-11-08T07:00:00Z",
			false, false, true,
		},
		{
			"2019-11-08T08:12:00Z", "3019-11-08T08:07:00Z",
			false, false, true,
		},
	} {
		t.Run(fmt.Sprintf("%d: %v %v", i, tc.t1, tc.t2), func(t *testing.T) {
			c1 := clock.FromTime(mustParse(tc.t1))
			c2 := clock.FromTime(mustParse(tc.t2))

			if c1 == c2 != tc.equals {
				t.Errorf("expected equals to be %v", tc.equals)
			}

			if c1.Before(c2) != tc.before {
				t.Errorf("expected before to be %v", tc.before)
			}

			if c1.After(c2) != tc.after {
				t.Errorf("expected after to be %v", tc.after)
			}
		})
	}

}

func TestClockString(t *testing.T) {
	for _, tc := range [...]struct {
		c    clock.Clock
		want string
	}{
		{
			clock.NewUTC(12, 10),
			"12:10",
		},
		{
			clock.NewUTC(23, 59),
			"23:59",
		},
		{
			clock.NewUTC(0, 0),
			"00:00",
		},
		{
			clock.NewUTC(12, 0),
			"12:00",
		},
		{
			clock.NewUTC(7, 12),
			"07:12",
		},
	} {
		t.Run(tc.want, func(t *testing.T) {
			if have := tc.c.String(); tc.want != have {
				t.Errorf("expected %q, have %q", tc.want, have)
			}
		})
	}
}

func TestClockUnmarshalJSON(t *testing.T) {
	for _, tc := range [...]struct {
		s    string
		want clock.Clock
	}{
		{
			"12:10",
			clock.NewUTC(12, 10),
		},
		{
			"23:59",
			clock.NewUTC(23, 59),
		},
		{
			"00:00",
			clock.NewUTC(0, 0),
		},
		{
			"12:00",
			clock.NewUTC(12, 0),
		},
		{
			"07:12",
			clock.NewUTC(7, 12),
		},
	} {
		t.Run(tc.s, func(t *testing.T) {
			var c clock.Clock
			err := c.UnmarshalJSON([]byte(tc.s))
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}

			if tc.want != c {
				t.Errorf("expected %s, have %s", tc.want, c)
			}
		})
	}
}
