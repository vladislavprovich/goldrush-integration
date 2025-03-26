package handler

import (
	"context"
	"errors"
	"github.com/vladislavprovich/goldrush-integration/internal/service"
	"net/http"
)

func (h *GoldRushHandler) GetTokenBalances(w http.ResponseWriter, r *http.Request) {
	var req service.TokenBalancesRequest
	h.handleJSONRequest(w, r, "GetTokenBalances", &req,
		func(ctx context.Context, reqObj interface{}) (interface{}, error) {
			typedReq, ok := reqObj.(*service.TokenBalancesRequest)
			if !ok {
				return nil, errors.New("handler.reqObj GetTokenBalances error")
			}
			return h.service.TokenBalances(ctx, typedReq)
		},
		".service.TokenBalances")
}
