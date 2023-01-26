package handler

import (
	"github.com/Egor-Tihonov/SandMailLogic/internal/service"
)

type Handler struct {
	se *service.Service
}

func New(ser *service.Service) *Handler {
	return &Handler{se: ser}
}