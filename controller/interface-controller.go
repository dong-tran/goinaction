package controller

import (
	"net/http"
)

type Controller interface {
	Path() string
	Handle(w http.ResponseWriter, r *http.Request)
}
