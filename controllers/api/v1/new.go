package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reanox/BitfinexLend/service/bitfinexService"
	"github.com/reanox/BitfinexLend/types"
)

func New(s *bitfinexService.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer c.JSON(http.StatusOK, types.Response{
			Message:   "ok",
			Status:    http.StatusOK,
			ErrorCode: 0,
		})
		apikey := c.PostForm("apikey")
		apisecret := c.PostForm("apisecret")
		s.CreateNewBFClient(apikey, apisecret)
	}
}
