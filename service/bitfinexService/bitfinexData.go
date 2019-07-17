package bitfinexService

import (
	"log"

	"github.com/bitfinexcom/bitfinex-api-go/v1"
)

var miniumLendNumber = 50.0
var annualizedRate30d = 75.0
var annualizedRateMin = 15.0

// A effective offer is 10 min
var offerRemoveTime int64 = 600

type BitfinexClient struct {
	APIKey       string
	APISecret    string
	Client       *bitfinex.Client
	MinRate      map[string]float64
	ReserveQuota map[string]float64
}

var BFClients []*BitfinexClient

func CreateNewBFClient(APIKey string, APISecret string) *BitfinexClient {
	c := &BitfinexClient{
		APIKey:       APIKey,
		APISecret:    APISecret,
		Client:       NewClient(APIKey, APISecret),
		MinRate:      make(map[string]float64),
		ReserveQuota: make(map[string]float64),
	}
	addBfcs(c)
	log.Printf("CreateNewBFClient: APIKey=%s", APIKey)
	return c
}

func removeBfcs(index int) {
	BFClients = append(BFClients[:index], BFClients[index+1:]...)
}

func addBfcs(bfc *BitfinexClient) {
	BFClients = append(BFClients, bfc)
}
