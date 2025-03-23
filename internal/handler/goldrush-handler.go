package handler

import (
	"encoding/json"
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
	log.Println("UserRegister called")         // Лог перед виконанням
	defer log.Println("UserRegister finished") // Лог після виконання

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
	h.logger.Info("handler.UserLogin called")
	var req service.LoginUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("handler.UserLogin.json.Decoder.Decode",
			slog.Any("error", err))
		http.Error(w, "json decode error", http.StatusBadRequest)
		return
	}

	token, err := h.service.LoginUser(r.Context(), &req)
	if err != nil {
		h.logger.Error("handler.UserLogin.service.Login",
			slog.Any("error", err))
		http.Error(w, "login user error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(token); err != nil {
		h.logger.Error("handler.UserLogin.json.Encoder.Encode",
			slog.Any("error", err))
		http.Error(w, "json encode error", http.StatusInternalServerError)
		return
	}
}

func (h *GoldRushHandler) GetTransactionInfo(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("handler.GetTransactionInfo called")
	var req service.TransactionInfoRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("handler.GetTransactionInfo.json.Decoder.Decode",
			slog.Any("error", err))
		http.Error(w, "json decode error", http.StatusBadRequest)
		return
	}

	transactionInfo, err := h.service.TransactionInfo(r.Context(), &req)
	if err != nil {
		h.logger.Error("handler.GetTransactionInfo.service.TransactionInfo",
			slog.Any("error", err))
		http.Error(w, "get transaction info error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(transactionInfo); err != nil {
		h.logger.Error("handler.GetTransactionInfo.json.Encoder.Encode",
			slog.Any("error", err))
		http.Error(w, "json encode error", http.StatusInternalServerError)
		return
	}
}

func (h *GoldRushHandler) GetTokenBalances(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("handler.GetTokenBalances called")
	var req service.TokenBalancesRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("handler.GetTokenBalances.json.Decoder.Decode",
			slog.Any("error", err))
		http.Error(w, "json decode error", http.StatusBadRequest)
		return
	}

	tokenBalances, err := h.service.TokenBalances(r.Context(), &req)
	if err != nil {
		h.logger.Error("handler.GetTokenBalances.service.TokenBalances",
			slog.Any("error", err))
		http.Error(w, "get token balances error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(tokenBalances); err != nil {
		h.logger.Error("handler.GetTokenBalances.json.Encoder.Encode",
			slog.Any("error", err))
		http.Error(w, "json encode error", http.StatusInternalServerError)
		return
	}
}

func (h *GoldRushHandler) GetRecentAddressTransaction(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("handler.GetRecentAddressTransaction called")
	var req service.RecentAddressTransactionRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("handler.GetRecentAddressTransaction.json.Decoder.Decode",
			slog.Any("error", err))
		http.Error(w, "json decode error", http.StatusBadRequest)
		return
	}

	addressTransaction, err := h.service.RecentAddressTransaction(r.Context(), &req)
	if err != nil {
		h.logger.Error("handler.GetRecentAddressTransaction.service.RecentAddressTransaction",
			slog.Any("error", err))
		http.Error(w, "get recent address transaction error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(addressTransaction); err != nil {
		h.logger.Error("handler.GetRecentAddressTransaction.json.Encoder.Encode",
			slog.Any("error", err))
		http.Error(w, "json encode error", http.StatusInternalServerError)
		return
	}
}

func (h *GoldRushHandler) GetHistoricalPortfolioTransaction(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("handler.GetHistoricalPortfolioTransaction called")
	var req service.HistoricalPortfolioValueRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		h.logger.Error("handler.GetHistoricalPortfolioTransaction.json.Decoder.Decode",
			slog.Any("error", err))
		http.Error(w, "json decode error", http.StatusBadRequest)
		return
	}

	historicalPortfolio, err := h.service.HistoricalPortfolioValue(r.Context(), &req)
	if err != nil {
		h.logger.Error("handler.GetHistoricalPortfolioTransaction.service.HistoricalPortfolioValue",
			slog.Any("error", err))
		http.Error(w, "get historical portfolio transaction error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err = json.NewEncoder(w).Encode(historicalPortfolio); err != nil {
		h.logger.Error("handler.GetHistoricalPortfolioTransaction.json.Encoder.Encode",
			slog.Any("error", err))
		http.Error(w, "json encode error", http.StatusInternalServerError)
		return
	}
}
