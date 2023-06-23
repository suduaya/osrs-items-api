package handler

import (
	"encoding/json"
	"net/http"
	"osrs-items-api/internal/items/service"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	// implicit swagger dep
	_ "osrs-items-api/pkg/oldschoolrs"
)

type Handler struct {
	svc *service.Service
}

func New(svc *service.Service) *Handler {
	return &Handler{
		svc: svc,
	}
}

// GetItemPriceVariationById Get Item by its Identifier
// ShowEntity godoc
// @Summary Provides an endpoint to get OSRS Item by ID
// @Description Endpoint to get OSRS Item by ID
// @Accept  json
// @Produce json
// @Param	id	path	int	true	"Item ID"
// @Success 200 {object} oldschoolrs.ItemPriceRecord "Runescape Item Price"
// @Failure 400 "Bad Request"
// @Failure 404 "Not Found"
// @Failure 500 "Internal Server Error"
// @Router /v1/items/{id}/price [get]
func (h *Handler) GetItemPriceVariationById(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	itemId := params["id"]

	itemIdInt, _ := strconv.Atoi(itemId)

	res, err := h.svc.GetItemPriceVariationById(itemIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			w.WriteHeader(http.StatusNotFound)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
