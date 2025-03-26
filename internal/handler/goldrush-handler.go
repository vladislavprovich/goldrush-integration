package handler

import (
	"context"
	"encoding/json"
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
