package middleware

import (
	"net/http"
)

type Secure struct{}

func (h Secure) Handle(handle func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" || len(r.Header["X"][0]) > 0 {
			handle(w, r)
		} else {
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte("Forbidden"))
		}
	}
}
