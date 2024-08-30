package types

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

/**
 * @example "2024-09-01"
 */
type CalendarDate string

const CalendarDateFormat = "2006-01-02"

// UnmarshalJSON implements json.Unmarshaler
func (cd *CalendarDate) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}

	*cd = CalendarDate(s)
	return nil
}

// MarshalJSON implements json.Marshaler
func (cd CalendarDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(cd))
}

// Scan implements sql.Scanner interface
func (cd *CalendarDate) Scan(value interface{}) error {
	if value == nil {
		*cd = ""
		return nil
	}
	switch v := value.(type) {
	case []byte:
		*cd = CalendarDate(string(v))
	case string:
		*cd = CalendarDate(v)
	case time.Time:
		*cd = CalendarDate(v.Format(CalendarDateFormat))
	default:
		return fmt.Errorf("cannot scan type %T into CalendarDate", value)
	}
	return nil
}

// Value implements driver.Valuer interface
func (cd CalendarDate) Value() (driver.Value, error) {
	return string(cd), nil
}

// Time converts CalendarDate to time.Time
func (cd CalendarDate) Time() (time.Time, error) {
	return time.Parse(CalendarDateFormat, string(cd))
}
