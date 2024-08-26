package handlers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/Wwreaker007/DIY-menu-service/common/data"
	"github.com/Wwreaker007/DIY-menu-service/kitchen/types"
)

type CreateOrderHttpHandler struct {
	kms 	types.Kitchen
}

func NewCreateOrderHandler(kms types.Kitchen) *CreateOrderHttpHandler {
	return &CreateOrderHttpHandler{
		kms: kms,
	}
}

func (h *CreateOrderHttpHandler) Handler(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithTimeout(r.Context(), time.Second*2)
	defer cancel()

	var request data.CreateOrderRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }
	response, _ := h.kms.CreateOrder(ctx, &request)
	json.NewEncoder(w).Encode(response)
}