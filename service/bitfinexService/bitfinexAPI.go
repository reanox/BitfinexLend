package bitfinexService

import (
	"fmt"
	"strconv"

	"github.com/bitfinexcom/bitfinex-api-go/v1"
)

func NewClient(APIKey string, APISecret string) *bitfinex.Client {
	client := bitfinex.NewClient().Auth(APIKey, APISecret)
	return client
}

func (bfc *BitfinexClient) GetLendBook(currency string, limitBids, limitAsks int) (bitfinex.Lendbook, error) {
	book, err := bfc.Client.Lendbook.Get(currency, limitBids, limitAsks)
	return book, err
}

func (bfc *BitfinexClient) Lend(currency string, limitBids, limitAsks int) (bitfinex.Lendbook, error) {
	book, err := bfc.Client.Lendbook.Get(currency, limitBids, limitAsks)
	return book, err
}

func (bfc *BitfinexClient) CreateLend(currency string, amount, rate float64, period int) (bitfinex.MarginOffer, error) {
	offer, err := bfc.Client.MarginFunding.NewLend(currency, amount, rate, period)
	return offer, err
}

func (bfc *BitfinexClient) GetFundingBalance() (float64, error) {
	balance, err := bfc.Client.Balances.All()
	for _, b := range balance {
		if b.Type == "deposit" && b.Currency == "usd" {
			available, _ := strconv.ParseFloat(b.Available, 64)
			return strconv.ParseFloat(fmt.Sprintf("%.2f", available), 64)
		}
	}
	return 0.0, err
}
