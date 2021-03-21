package clock

import (
	"time"
)

const layout = "15:04 MST"

// Clock holds information for a specific hour of the day, in a specific
// timezone.
type Clock struct {
	hour     int
	minute   int
	location *time.Location
}

func (c Clock) Hour() int {
	return c.hour
}

func (c Clock) Minute() int {
	return c.minute
}

// String implements fmt.Stringer
func (c Clock) String() string {
	return time.Date(1, 1, 1, c.hour, c.minute, 0, 0, c.location).Format(layout)
}

// MarshalJSON implements json.Marshaler
func (c Clock) MarshalJSON() ([]byte, error) {
	return []byte(`"` + c.String() + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler. The time is expected to be a
// quoted string in "15:04 MST" format.
func (c *Clock) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}

	t, err := time.Parse(`"`+layout+`"`, string(data))
	*c = FromTime(t)
	return err
}

func (c Clock) In(location *time.Location) Clock {
	return FromTime(c.time().In(location))
}

func (c Clock) time() time.Time {
	return time.Date(1, 1, 1, c.hour, c.minute, 0, 0, c.location)
}

func New(hour, minute int, location *time.Location) Clock {
	return FromTime(time.Date(1, 1, 1, hour, minute, 0, 0, location))
}

func NewUTC(hour, minute int) Clock {
	return New(hour, minute, time.UTC)
}

func FromTime(t time.Time) Clock {
	return Clock{
		hour:     t.Hour(),
		minute:   t.Minute(),
		location: t.Location(),
	}
}

func (c Clock) Before(u Clock) bool {
	if c.Hour() < u.Hour() {
		return true
	}

	if c.Hour() > u.Hour() {
		return false
	}

	return c.Minute() < u.Minute()
}

func (c Clock) After(u Clock) bool {
	if c.Hour() > u.Hour() {
		return true
	}

	if c.Hour() < u.Hour() {
		return false
	}

	return c.Minute() > u.Minute()
}
