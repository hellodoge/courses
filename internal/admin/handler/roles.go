package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hellodoge/courses-tg-bot/courses/messages"
	"net/http"
)

func (h *Handler) initRolesRoutes(group gin.IRouter) {
	group.POST("/new-moderator", h.newModerator)
}

type newModeratorQuery struct {
	Description string `json:"description"`
}

type newModeratorResponse struct {
	Token string `json:"token"`
}

func (h *Handler) newModerator(ctx *gin.Context) {
	if !h.userIsAdmin(ctx) {
		errorResponse(ctx, messages.NotAnAdministrator, http.StatusForbidden)
		return
	}
	var input newModeratorQuery
	if err := ctx.BindJSON(&input); err != nil {
		errorResponse(ctx, err.Error(), http.StatusBadRequest)
		return
	}
	token, err := h.services.Roles.NewModerator(input.Description)
	if err != nil {
		errorResponse(ctx, unknownError, http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, newModeratorResponse{Token: token})
}
