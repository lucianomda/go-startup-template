package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) HandlePing(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
