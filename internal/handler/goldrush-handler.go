package handler

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"log/slog"
	"net/http"

	"github.com/vladislavprovich/goldrush-integration/internal/service"
)

type GoldRushHandler struct {
	service service.Service
	logger  *slog.Logger
	cfg     *Config
}

func NewGoldRushHandler(srv *service.Service, log *slog.Logger, cfg *Config) *GoldRushHandler {
	return &GoldRushHandler{
		service: *srv,
		logger:  log,
		cfg:     cfg,
	}
}

func (h *GoldRushHandler) UserRegister(w http.ResponseWriter, r *http.Request) {
	log.Println("UserRegister called")
	defer log.Println("UserRegister finished")

	h.logger.Info("handler.UserRegister called")
	var req service.RegisterUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("handler.UserRegister.json.Decoder.Decode",
			slog.Any("error", err))
		http.Error(w, "json decode error", http.StatusBadRequest)
		return
	}

	userID, err := h.service.RegisterUser(r.Context(), &req)
	if err != nil {
		h.logger.Error("handler.UserRegister.service.Register",
			slog.Any("error", err))
		http.Error(w, "register user error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	if err = json.NewEncoder(w).Encode(userID); err != nil {
		h.logger.Error("handler.UserRegister.json.Encoder.Encode",
			slog.Any("error", err))
		http.Error(w, "json encode error", http.StatusInternalServerError)
		return
	}
}

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

// handleJSONRequest is a generic handler helper function to reduce code duplication.
// It handles the common pattern of decoding a JSON request, calling a service method, and encoding a JSON response.
func (h *GoldRushHandler) handleJSONRequest(
	w http.ResponseWriter,
	r *http.Request,
	handlerName string,
	requestObj interface{},
	serviceFunc func(
		context.Context,
		interface{},
	) (interface{}, error), serviceErrMsg string) {
	h.logger.Info("handler." + handlerName + " called")

	if err := json.NewDecoder(r.Body).Decode(requestObj); err != nil {
		h.logger.Error("handler."+handlerName+".json.Decoder.Decode",
			slog.Any("error", err))
		http.Error(w, "json decode error", http.StatusBadRequest)
		return
	}

	response, err := serviceFunc(r.Context(), requestObj)
	if err != nil {
		h.logger.Error("handler."+handlerName+serviceErrMsg,
			slog.Any("error", err))
		http.Error(w, serviceErrMsg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(response); err != nil {
		h.logger.Error("handler."+handlerName+".json.Encoder.Encode",
			slog.Any("error", err))
		http.Error(w, "json encode error", http.StatusInternalServerError)
		return
	}
}

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
