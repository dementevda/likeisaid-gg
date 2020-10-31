package models

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

type JsonTimestamp time.Time

func (j *JsonTimestamp) UnmarshalJSON(b []byte) error {
	s := strings.Trim(string(b), "\"")
	t, err := time.Parse(time.RFC3339, s)
	fmt.Println(t)
	if err != nil {
		return err
	}
	*j = JsonTimestamp(t)
	return nil
}

func (j JsonTimestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(j)
}

func (j JsonTimestamp) Format(s string) string {
	t := time.Time(j)
	return t.Format(s)
}
