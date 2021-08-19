package customTypes

import (
	"bytes"
	"database/sql/driver"
	"fmt"
	"time"
)

type YmdTime struct {
	time.Time
}

const ymdLayout = "2006-01-02"

func (yt *YmdTime) UnmarshalJSON(value []byte) (err error) {
	value = bytes.Trim(value, "\"")
	yt.Time, err = time.Parse(ymdLayout, string(value))

	return
}

func (yt YmdTime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", yt.Format(ymdLayout))), nil
}

func (yt *YmdTime) Scan(src interface{}) error {
	if t, ok := src.(time.Time); ok {
		yt.Time = t
	}

	return nil
}

func (yt YmdTime) Value() (driver.Value, error) {
	return yt.Time, nil
}
