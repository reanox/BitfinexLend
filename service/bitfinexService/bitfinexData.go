package bitfinexService

import (
	"github.com/bitfinexcom/bitfinex-api-go/v1"
)

var BFClients []struct {
	APIKey    string
	APISecret string
	Client    *bitfinex.Client
}
