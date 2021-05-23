package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hellodoge/courses-tg-bot/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	var router = gin.New()
	api := router.Group("/api", h.userIdentity)
	{
		roles := api.Group("/roles")
		h.initRolesRoutes(roles)
	}
	return router
}
