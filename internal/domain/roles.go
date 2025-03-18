package domain

import (
	"database/sql/driver"
	"errors"
)

type Role string

const (
	Admin     Role = "admin"
	Moderator Role = "moderator"
	Client    Role = "user"
	Guest     Role = "guest"
)

func (r Role) String() string {
    return string(r)
}

func (r *Role) Scan(value any) error {
	str, ok := value.(string)
	if !ok {
		return errors.New("invalid type for Role")
	}
	*r = Role(str)
	return nil
}

func (r Role) Value() (driver.Value, error) {
	return r.String(), nil
}
