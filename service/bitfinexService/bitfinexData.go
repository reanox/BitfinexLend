package bitfinexService

import (
	"github.com/bitfinexcom/bitfinex-api-go/v1"
)

var miniumLendNumber = 50.0
var AnnualizedRate30d = 75.0
var AnnualizedRateMin = 15.0

type BitfinexClient struct {
	APIKey       string
	APISecret    string
	Client       *bitfinex.Client
	MinRate      map[string]float64
	ReserveQuota map[string]float64
}

var BFClients []*BitfinexClient

func CreateNewBFClient(APIKey string, APISecret string) *BitfinexClient {
	return &BitfinexClient{
		APIKey:       APIKey,
		APISecret:    APISecret,
		Client:       NewClient(APIKey, APISecret),
		MinRate:      make(map[string]float64),
		ReserveQuota: make(map[string]float64),
	}
}

func removeBfcs(index int) {
	BFClients = append(BFClients[:index], BFClients[index+1:]...)
}

func addBfcs(bfc *BitfinexClient) {
	BFClients = append(BFClients, bfc)
}
