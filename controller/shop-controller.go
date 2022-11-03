package controller

import (
	"encoding/json"
	"net/http"

	"github.com/dong-tran/goinaction/utils"
)

type ShopController struct {
	utils.AccessLog
}

func (c ShopController) Path() string {
	return "/shops"
}

func (c ShopController) Handle(w http.ResponseWriter, r *http.Request) {
	c.LogAccess(r)
	var s = shop{
		ID:      1,
		Name:    "Test shop",
		Address: "20 NTMK",
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(s)
}

type shop struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}
