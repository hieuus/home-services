package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
)

func (u *User) Scan(value interface{}) (err error) {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New(fmt.Sprint("failed to unmarshal User value: ", value))
	}
	return json.Unmarshal(bytes, &u)
}

func (u *User) Value() (driver.Value, error) {
	return json.Marshal(u)
}
