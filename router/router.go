package router

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	v1 "github.com/reanox/BitfinexLend/controllers/api/v1"
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

	APIv1 := router.Group("/api/v1")
	{
		APIv1.POST("/new", v1.New)
	}
	return router
}
