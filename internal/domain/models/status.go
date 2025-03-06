package models

import (
	"database/sql/driver"
	"errors"
)

type Status string

const (
	Accepted Status  = "accepted"
	Pending  Status  = "pending"
	Rejected Status = "rejected"
)

func (s Status) String() string {
	return string(s)
}

func (s *Status) Scan(value any) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("invalid type for Role")
	}
	*s = Status(str)
	return nil
}

func (s Status) Value() (driver.Value, error) {
	return s.String(), nil
}
