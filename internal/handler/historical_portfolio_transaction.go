package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/vladislavprovich/goldrush-integration/internal/service"
)

func (h *GoldRushHandler) GetHistoricalPortfolioTransaction(w http.ResponseWriter, r *http.Request) {
	var req service.HistoricalPortfolioValueRequest
	h.handleJSONRequest(w, r, "GetHistoricalPortfolioTransaction", &req,
		func(ctx context.Context, reqObj interface{}) (interface{}, error) {
			typedReq, ok := reqObj.(*service.HistoricalPortfolioValueRequest)
			if !ok {
				return nil, errors.New("handler.reqObj GetHistoricalPortfolioTransaction error")
			}
			return h.service.HistoricalPortfolioValue(ctx, typedReq)
		},
		".service.HistoricalPortfolioValue")
}
