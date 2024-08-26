package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Wwreaker007/DIY-menu-service/common/data"
	"github.com/Wwreaker007/DIY-menu-service/kitchen/types"
)

type GetOrderHttpHandler struct {
	kms		types.Kitchen
}

func NewGetOrderHandler(kms types.Kitchen) *GetOrderHttpHandler {
	return &GetOrderHttpHandler{
		kms: kms,
	}
}

func (h *GetOrderHttpHandler) Handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	var request data.GetOrderRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	response, _ := h.kms.GetOrder(ctx, &request)
	json.NewEncoder(w).Encode(response)
}