package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dong-tran/goinaction/dto"
	"github.com/dong-tran/goinaction/service"
	"github.com/dong-tran/goinaction/utils"
)

type DummyController struct {
	utils.AccessLog
}

func (c DummyController) Path() string {
	return "/dummies"
}

func (c DummyController) Handle(w http.ResponseWriter, r *http.Request) {
	c.LogAccess(r)
	w.Header().Add("Content-Type", "application/json")
	if r.Method == "GET" {
		c.handleGet(w, r)
	} else if r.Method == "POST" {
		c.handlePost(w, r)
	}
}

func (c DummyController) handleGet(w http.ResponseWriter, r *http.Request) {
	var sv = service.DummyService{}
	all, err := sv.GetAll()
	if err != nil {
		log.Printf("ERROR\t%+v", err)
	} else {
		json.NewEncoder(w).Encode(all)
	}
}

func (c DummyController) handlePost(w http.ResponseWriter, r *http.Request) {
	var dto dto.Dummy
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&dto)
	var sv = service.DummyService{}
	err := sv.Insert(&dto)
	if err != nil {
		log.Printf("ERROR\t%+v", err)
	} else {
		json.NewEncoder(w).Encode("{\"status\":\"OK\"}")
	}

}
