package views

import (
	"errors"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
)

func isDup(err error) bool {
	var e mongo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}
	return false
}

func validDate(date time.Time) error {
	if date.Before(time.Now()) {
		return errors.New("Date should be greater then now")
	}
	return nil
}
