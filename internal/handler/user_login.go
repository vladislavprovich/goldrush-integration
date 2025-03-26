package handler

import (
	"context"
	"errors"
	"github.com/vladislavprovich/goldrush-integration/internal/service"
	"net/http"
)

func (h *GoldRushHandler) UserLogin(w http.ResponseWriter, r *http.Request) {
	var req service.LoginUserRequest
	h.handleJSONRequest(w, r, "UserLogin", &req,
		func(ctx context.Context, reqObj interface{}) (interface{}, error) {
			typedReq, ok := reqObj.(*service.LoginUserRequest)
			if !ok {
				return nil, errors.New("handler.reqObj UserLogin error")
			}
			return h.service.LoginUser(ctx, typedReq)
		},
		".service.Login")
}
