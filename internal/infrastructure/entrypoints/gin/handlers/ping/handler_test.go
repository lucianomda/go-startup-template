package ping

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/lucianomda/go-startup-template/internal/infrastructure/entrypoints/gin/httptests"
	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	h := NewHandler()
	assert.IsType(t, &Handler{}, h)
}

func givenHandlePingRequestContext() (*gin.Context, *httptest.ResponseRecorder) {
	return httptests.GetGinContextStub("", http.MethodGet, "ping", nil, nil, nil)
}

func TestHandler_HandlePing(t *testing.T) {
	h := NewHandler()
	ginCtx, responseRecorder := givenHandlePingRequestContext()
	h.HandlePing(ginCtx)

	assert.Equal(t, http.StatusOK, responseRecorder.Code)
	assert.Equal(t, "pong", responseRecorder.Body.String())
}
