package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/reanox/BitfinexLend/service/bitfinexService"
)

func New(c *gin.Context) {
	defer c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
	apikey := c.PostForm("apikey")
	apisecret := c.PostForm("apisecret")
	bitfinexService.CreateNewBFClient(apikey, apisecret)
}
