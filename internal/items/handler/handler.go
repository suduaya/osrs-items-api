package handler

import (
	"encoding/json"
	"net/http"
	"osrs-items-api/internal/items/service"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

type Handler struct {
	svc *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

func (h *Handler) GetItemById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemId := params["id"]

	itemIdInt, _ := strconv.Atoi(itemId)

	if res, err := h.svc.FindItemById(itemIdInt); err != nil {
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
