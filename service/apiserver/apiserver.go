package apiserver

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	serviceconfig "github.com/reanox/BitfinexLend/config"
	v1 "github.com/reanox/BitfinexLend/controllers/api/v1"
	"github.com/reanox/BitfinexLend/service/bitfinexService"
)

func defaultRouter() *gin.Engine {
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

// New a apiserver
func New(s *bitfinexService.Service) *http.Server {

	router := defaultRouter()
	APIv1 := router.Group("/api/v1")
	{
		APIv1.POST("/new", v1.New(s))
	}

	server := &http.Server{
		Addr:    serviceconfig.BindAddr,
		Handler: router,
	}
	return server
}
