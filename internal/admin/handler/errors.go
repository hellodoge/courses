package handler

import "github.com/gin-gonic/gin"

const (
	emptyAuthorizationHeader = "Empty '" + authorizationHeader + "' header"
	unknownError             = "Unknown Internal Server Error. Please, contact API service team"
)

func errorResponse(c *gin.Context, msg string, statusCode int) {
	c.AbortWithStatusJSON(statusCode, struct {
		Error string `json:"error"`
	}{
		Error: msg,
	})
}
