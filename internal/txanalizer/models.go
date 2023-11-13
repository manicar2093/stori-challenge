package txanalizer

import (
	"fmt"
	"strings"
	"time"
)

type (
	Transaction struct {
		Id     uint
		Date   Date
		Amount float64
	}

	TransactionsAnalizys struct {
		TotalBalance       float64
		TransactionByMonth struct {
			Month    string
			Quantity uint
		}
		AverageDebitAmount  float64
		AverageCreditAmount float64
	}

	Date time.Time
)

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

func (c Date) Day() int {
	return time.Time(c).Day()
}

func (c *Date) unmarshal(data []byte) (err error) {
	s := strings.Trim(string(data), "\"")
	if s == "null" {
		return
	}
	parsedTime, err := time.Parse(time.DateOnly, s)
	if err != nil {
		return
	}
	*c = Date(parsedTime)
	return
}

func (c Date) marshal() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", time.Time(c).Format(time.DateOnly))), nil
}
