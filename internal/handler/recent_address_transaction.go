package handler

import (
	"context"
	"errors"
	"github.com/vladislavprovich/goldrush-integration/internal/service"
	"net/http"
)

func (h *GoldRushHandler) GetRecentAddressTransaction(w http.ResponseWriter, r *http.Request) {
	var req service.RecentAddressTransactionRequest
	h.handleJSONRequest(w, r, "GetRecentAddressTransaction", &req,
		func(ctx context.Context, reqObj interface{}) (interface{}, error) {
			typedReq, ok := reqObj.(*service.RecentAddressTransactionRequest)
			if !ok {
				return nil, errors.New("handler.reqObj GetRecentAddressTransaction error")
			}
			return h.service.RecentAddressTransaction(ctx, typedReq)
		},
		".service.RecentAddressTransaction")
}
