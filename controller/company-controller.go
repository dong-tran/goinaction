package controller

import (
	"fmt"
	"net/http"

	"github.com/dong-tran/goinaction/utils"
)

type CompanyController struct {
	utils.AccessLog
}

func (c CompanyController) Path() string {
	return "/companies"
}

func (c CompanyController) Handle(w http.ResponseWriter, r *http.Request) {
	c.LogAccess(r)
	fmt.Fprintf(w, "Welcome to the Company page!")
}
