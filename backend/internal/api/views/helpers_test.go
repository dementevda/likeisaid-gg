package views

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/mongo"
)

type testCaseIsDup struct {
	err    error
	result bool
}

func TestIsDup(t *testing.T) {
	cases := []testCaseIsDup{
		testCaseIsDup{
			err:    fmt.Errorf("%s", "err"),
			result: false,
		},
		testCaseIsDup{
			err: mongo.WriteException{
				WriteErrors: []mongo.WriteError{
					mongo.WriteError{
						Code: 200,
					},
				},
			},
			result: false,
		},
		testCaseIsDup{
			err: mongo.WriteException{
				WriteErrors: []mongo.WriteError{
					mongo.WriteError{
						Code: 200,
					},
					mongo.WriteError{
						Code: 200,
					},
				},
			},
			result: false,
		},
		testCaseIsDup{
			err: mongo.WriteException{
				WriteErrors: []mongo.WriteError{
					mongo.WriteError{
						Code: 200,
					},
					mongo.WriteError{
						Code: 11000,
					},
				},
			},
			result: true,
		},
	}

	for _, tCase := range cases {
		assert.Equal(t, isDup(tCase.err), tCase.result)
	}
}

type testCaseValidDate struct {
	timestamp time.Time
	isErr     bool
	name      string
}

func TestValidDate(t *testing.T) {
	cases := []testCaseValidDate{
		testCaseValidDate{
			timestamp: time.Now().Add(-time.Hour * 24),
			isErr:     true,
			name:      "Lower then now",
		},
		testCaseValidDate{
			timestamp: time.Now().Add(time.Hour * 24),
			isErr:     false,
			name:      "Greater then now",
		},
	}

	for _, tCase := range cases {
		err := validDate(tCase.timestamp)
		if err == nil && tCase.isErr {
			t.Errorf("%s", tCase.name)
		}

	}
}
