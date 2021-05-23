package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
)

func (h *Handler) userIsAdmin(ctx *gin.Context) bool {
	token, err := getToken(ctx)
	if err != nil {
		errorResponse(ctx, unknownError, http.StatusInternalServerError)
		logrus.Error(err)
		return false
	}
	isAdmin, err := h.services.Roles.UserIsAdmin(token)
	if err != nil {
		errorResponse(ctx, unknownError, http.StatusInternalServerError)
		logrus.Error(err)
		return false
	}
	return isAdmin
}

func (h *Handler) userIsModerator(ctx *gin.Context) bool {
	token, err := getToken(ctx)
	if err != nil {
		errorResponse(ctx, unknownError, http.StatusInternalServerError)
		logrus.Error(err)
		return false
	}
	isModerator, err := h.services.Roles.UserIsModerator(token)
	if err != nil {
		errorResponse(ctx, unknownError, http.StatusInternalServerError)
		logrus.Error(err)
		return false
	}
	return isModerator
}
