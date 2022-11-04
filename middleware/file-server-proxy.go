package middleware

import (
	"net/http"

	"github.com/dong-tran/goinaction/utils"
)

type Proxy struct {
	handler http.Handler
	utils.AccessLog
}

func (p Proxy) Handler(handler http.Handler) http.Handler {
	return &Proxy{handler: handler}
}

func (p Proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	p.LogAccess(r)
	p.handler.ServeHTTP(w, r)
}
