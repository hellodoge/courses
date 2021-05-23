package handler

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

const (
	authorizationHeader = "Authorization"
	tokenKey            = "token"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		errorResponse(ctx, emptyAuthorizationHeader, http.StatusUnauthorized)
		return
	}

	ctx.Set(tokenKey, strings.TrimPrefix(header, "Bearer "))
}

func getToken(ctx *gin.Context) (string, error) {
	tokenContext, ok := ctx.Get(tokenKey)
	if !ok {
		return "", errors.New("token not found")
	}

	token, ok2 := tokenContext.(string)
	if !ok2 {
		return "", fmt.Errorf("token has other type, than string (%T)", tokenContext)
	}
	return token, nil
}
