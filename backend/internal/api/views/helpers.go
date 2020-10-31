package views

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
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

func writeError(w http.ResponseWriter, errStatus int, apiErr interface{}) {
	w.WriteHeader(errStatus)
	json.NewEncoder(w).Encode(&apiErr)
	fmt.Fprintln(w)
}

func writeResponse(w http.ResponseWriter, status int, data interface{}) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
	fmt.Fprintln(w)
}
