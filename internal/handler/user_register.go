package handler

import (
	"encoding/json"
	"github.com/vladislavprovich/goldrush-integration/internal/service"
	"log"
	"log/slog"
	"net/http"
)

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
