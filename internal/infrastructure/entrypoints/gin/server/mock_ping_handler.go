// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package server

import (
	gin "github.com/gin-gonic/gin"
	mock "github.com/stretchr/testify/mock"
)

// MockPingHandler is an autogenerated mock type for the PingHandler type
type MockPingHandler struct {
	mock.Mock
}

// HandlePing provides a mock function with given fields: ctx
func (_m *MockPingHandler) HandlePing(ctx *gin.Context) {
	_m.Called(ctx)
}