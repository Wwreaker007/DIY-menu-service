package handlers

import (
	"time"
	"context"
	"net/http"
	"encoding/json"

	"github.com/Wwreaker007/DIY-menu-service/cookhouse/types"
	"github.com/Wwreaker007/DIY-menu-service/common/data"

)

type GetOrderByStatusFilterHandler struct {
	cms 	types.CookHouse
}

func NewGetAllOrderByStatusFilterHandler(cms types.CookHouse) *GetOrderByStatusFilterHandler {
	return &GetOrderByStatusFilterHandler{
		cms: cms,
	}
}

func (h *GetOrderByStatusFilterHandler) Handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	var request data.GetOrderByStatusFilterRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	response, err := h.cms.GetAllOrdersByStatusFilter(ctx, &request)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	json.NewEncoder(w).Encode(response)
}