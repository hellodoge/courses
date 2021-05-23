package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/hellodoge/courses-tg-bot/courses"
	"github.com/hellodoge/courses-tg-bot/courses/messages"
	"net/http"
)

func (h *Handler) initCoursesRoutes(group gin.IRouter) {
	group.POST("/new", h.newCourse)
}

type newCourseResponse struct {
	ID string `json:"id"`
}

func (h *Handler) newCourse(ctx *gin.Context) {
	if !h.userIsModerator(ctx) {
		errorResponse(ctx, messages.NotAModerator, http.StatusForbidden)
		return
	}
	var input courses.Course
	if err := ctx.BindJSON(&input); err != nil {
		errorResponse(ctx, err.Error(), http.StatusBadRequest)
		return
	}
	id, err := h.services.NewCourse(&input)
	if err != nil {
		errorResponse(ctx, unknownError, http.StatusInternalServerError)
		return
	}
	ctx.JSON(http.StatusOK, newCourseResponse{ID: id})
}
