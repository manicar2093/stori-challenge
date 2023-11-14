package txanalizer

import (
	"database/sql/driver"
	"fmt"
	"strconv"
	"strings"
	"time"
)

const dateFormat = "1/2"

type Date struct {
	day   int
	month time.Month
}

func NewDate(month time.Month, day int) Date {
	d := time.Date(0, time.Month(month), day, 0, 0, 0, 0, time.Local)
	return Date{
		day:   d.Day(),
		month: d.Month(),
	}
}

func (c *Date) UnmarshalJSON(data []byte) error {
	return c.unmarshal(data)
}

func (c Date) MarshalJSON() ([]byte, error) {
	return c.marshal()
}

func (c Date) MarshalCSV() ([]byte, error) {
	return c.marshal()
}

func (c *Date) UnmarshalCSV(data []byte) error {
	return c.unmarshal(data)
}

func (c *Date) Scan(src any) error {
	asString := src.(string)
	splitted := strings.Split(asString, "/")
	month, err := strconv.Atoi(splitted[0])
	if err != nil {
		return err
	}
	day, err := strconv.Atoi(splitted[1])
	if err != nil {
		return err
	}
	c.month = time.Month(month)
	c.day = day
	return nil
}

func (c Date) Value() (driver.Value, error) {
	return fmt.Sprintf("%d/%d", c.month, c.day), nil
}

func (c Date) Day() int {
	return c.day
}

func (c Date) Month() time.Month {
	return c.month
}

func (c *Date) unmarshal(data []byte) (err error) {
	s := strings.Trim(string(data), "\"")
	if s == "null" {
		return
	}
	parsedTime, err := time.Parse(dateFormat, s)
	if err != nil {
		return
	}
	*c = Date{
		day:   parsedTime.Day(),
		month: parsedTime.Month(),
	}
	return
}

func (c Date) marshal() ([]byte, error) {
	return []byte(fmt.Sprintf("%d/%d", c.month, c.day)), nil
}
