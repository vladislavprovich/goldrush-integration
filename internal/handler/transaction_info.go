package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/vladislavprovich/goldrush-integration/internal/service"
)

func (h *GoldRushHandler) GetTransactionInfo(w http.ResponseWriter, r *http.Request) {
	var req service.TransactionInfoRequest
	h.handleJSONRequest(w, r, "GetTransactionInfo", &req,
		func(ctx context.Context, reqObj interface{}) (interface{}, error) {
			typedReq, ok := reqObj.(*service.TransactionInfoRequest)
			if !ok {
				return nil, errors.New("handler.reqObj GetTransactionInfo error")
			}
			return h.service.TransactionInfo(ctx, typedReq)
		},
		".service.TransactionInfo")
}
