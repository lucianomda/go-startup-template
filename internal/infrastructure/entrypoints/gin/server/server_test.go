package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNew_whenPassingPingHandlerThenItsPathIsConfigured(t *testing.T) {
	method := http.MethodGet
	path := "/ping"
	expectedResponseStatus := http.StatusOK
	request, _ := http.NewRequest(method, path, nil)
	responseRecorder := httptest.NewRecorder()

	mockPingHandler := new(MockPingHandler)
	mockPingHandler.On("HandlePing", mock.AnythingOfType("*gin.Context")).Return()

	s := New(mockPingHandler)
	s.ServeHTTP(responseRecorder, request)

	assert.Equal(t, expectedResponseStatus, responseRecorder.Code)
	mockPingHandler.AssertExpectations(t)
}
