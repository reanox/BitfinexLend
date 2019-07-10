package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

// New router
func New() *gin.Engine {
	router := gin.Default()
	config := cors.DefaultConfig()
	config.AllowCredentials = true
	config.AllowOriginFunc = func(origin string) bool {
		return true
	}
	router.Use(cors.New(config))
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Server is running")
	})

	return router
}
