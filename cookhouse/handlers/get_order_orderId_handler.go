package handlers

import (
	"time"
	"context"
	"net/http"
	"encoding/json"

	"github.com/Wwreaker007/DIY-menu-service/cookhouse/types"
	"github.com/Wwreaker007/DIY-menu-service/common/data"

)

type GetOrderByOrderIDHandler struct {
	cms 	types.CookHouse
}

func NewGetOrderByOrderIDHandler(cms types.CookHouse) *GetOrderByOrderIDHandler {
	return &GetOrderByOrderIDHandler{
		cms: cms,
	}
}

func (h *GetOrderByOrderIDHandler) Handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	var request data.GetOrderByOrderIDRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	response, _ := h.cms.GetOrderByOrderID(ctx, &request)
	json.NewEncoder(w).Encode(response)
}