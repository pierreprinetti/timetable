package clock

import (
	"time"
)

type Clock struct {
	hour     int
	minute   int
	location *time.Location
}

type clockError string

func (cErr clockError) Error() string {
	return string(cErr)
}

const (
	ErrHourOutOfBounds   clockError = "The hour must be between 0 and 23"
	ErrMinuteOutOfBounds clockError = "The minute must be between 0 and 59"
)

func (c Clock) Hour() int {
	return c.hour
}

func (c Clock) Minute() int {
	return c.minute
}

// String implements fmt.Stringer
func (c Clock) String() string {
	return time.Date(1, 1, 1, c.hour, c.minute, 0, 0, c.location).Format("15:04")
}

// MarshalJSON implements json.Marshaler
func (c Clock) MarshalJSON() ([]byte, error) {
	return []byte(`"` + c.String() + `"`), nil
}

// UnmarshalJSON implements json.Unmarshaler
func (c *Clock) UnmarshalJSON(src []byte) error {
	t, err := time.Parse("15:04", string(src))
	if err != nil {
		return err
	}

	*c = FromTime(t)

	return nil
}

func (c Clock) IsZero() bool {
	return c.Hour() == 0 && c.Minute() == 0
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
