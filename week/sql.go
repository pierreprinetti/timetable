package week

import (
	"database/sql/driver"
	"fmt"
)

func (w *Week) Scan(src interface{}) error {
	var scanned Week

	raw, ok := src.(string)
	if !ok {
		return ErrNotString
	}

	if len(raw) == 0 {
		return ErrEmptyString
	}

	if _, err := fmt.Sscanf(raw, "%b", &scanned); err != nil {
		return err
	}

	*w = scanned
	return nil
}

func (w Week) Value() (driver.Value, error) {
	return fmt.Sprintf("%b", w), nil
}
