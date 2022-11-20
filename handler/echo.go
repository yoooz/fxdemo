package handler

import (
	"io"
	"net/http"

	"github.com/yoooz/fxdemo/service"
	"go.uber.org/zap"
)

type EchoHandler struct {
	log *zap.Logger

	service *service.EchoService
}

func newEchoHandler(log *zap.Logger, service *service.EchoService) *EchoHandler {
	return &EchoHandler{
		log:     log,
		service: service,
	}
}

func (*EchoHandler) Pattern() string {
	return "/echo"
}

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if _, err := io.Copy(w, r.Body); err != nil {
		h.log.Warn("Failed to handle request", zap.Error(err))
	}

	h.service.CallService()
}
