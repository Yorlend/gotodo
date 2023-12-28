package handlers

import (
	"net/http"

	"go.uber.org/zap"
)

type HomeHandler struct {
	logger *zap.SugaredLogger
}

func NewHomeHandler(
	logger *zap.SugaredLogger,
) *HomeHandler {
	return &HomeHandler{
		logger: logger,
	}
}

func (h *HomeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.logger.Info("Serving homepage")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}
