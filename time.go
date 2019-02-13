package tradesatoshi

import (
	"encoding/json"
	"time"
)

const TIME_FORMAT = "2006-01-02T15:04:05"

type Time struct {
	time.Time
}

func (ft *Time) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse(TIME_FORMAT, s)
	if err != nil {
		return err
	}
	ft.Time = t
	return nil
}

func (ft Time) MarshalJSON() ([]byte, error) {
	return json.Marshal((*time.Time)(&ft.Time).Format(TIME_FORMAT))
}
