package views

import (
	"io"
	"net/http"
)

// HandleHello returns "Hello, World!"
func HandleHello() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello, World!")
	}
}
