package handlers

import (
	"time"
	"context"
	"net/http"
	"encoding/json"

	"github.com/Wwreaker007/DIY-menu-service/cookhouse/types"
	"github.com/Wwreaker007/DIY-menu-service/common/data"

)

type UpdateOrderStatusHandler struct {
	cms 	types.CookHouse
}

func NewUpdateOrderStatusHandler(cms types.CookHouse) *UpdateOrderStatusHandler {
	return &UpdateOrderStatusHandler{
		cms: cms,
	}
}

func (h *UpdateOrderStatusHandler) Handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	var request data.UpdateOrderStatusRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	response, _ := h.cms.UpdateOrderStatus(ctx, &request)
	json.NewEncoder(w).Encode(response)
}