package utils

import (
	"log"
	"net/http"
)

type AccessLog struct{}

func (a AccessLog) LogAccess(r *http.Request) {
	log.Printf("%s\t%s\t%s", r.Method, r.RequestURI, r.Header["User-Agent"][0])
}
