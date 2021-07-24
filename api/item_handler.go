package api

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
)

func (t *API) GetItemById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemId := params["id"]

	if res, err := t.ItemManager.FindItemById(itemId); err != nil {
		if strings.Contains(err.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}

func (t *API) GetItemByName(w http.ResponseWriter, r *http.Request) {
	itemName := r.FormValue("name")

	if res, err := t.ItemManager.FindItemByName(itemName); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(res)
	}
}
