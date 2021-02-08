package server

import (
	"fmt"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
)

type PingHandler interface {
	HandlePing(ctx *gin.Context)
}

type routerConfig struct {
	EnableResponseCompressionSupport bool
}

func customRouter(conf routerConfig) *gin.Engine {
	router := gin.New()

	if conf.EnableResponseCompressionSupport {
		router.Use(gzip.Gzip(gzip.DefaultCompression))
	}

	router.NoRoute(noRouteHandler)
	return router
}

func noRouteHandler(c *gin.Context) {
	c.JSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Resource %s not found.", c.Request.URL.Path)})
}

func New(pingHandler PingHandler) *gin.Engine {
	config := routerConfig{
		EnableResponseCompressionSupport: true,
	}
	router := customRouter(config)
	configHandlers(router, pingHandler)

	return router
}

func configHandlers(router gin.IRouter, pingHandler PingHandler) {
	router.GET("ping", pingHandler.HandlePing)
}
