package bitfinexService

import (
	"fmt"
	"strconv"

	"github.com/bitfinexcom/bitfinex-api-go/v1"
	"github.com/reanox/BitfinexLend/db/models"
)

// BitfinexClient is setting by user
type BitfinexClient struct {
	Client *bitfinex.Client
	User   *models.User
}

func NewClient(APIKey string, APISecret string) *bitfinex.Client {
	client := bitfinex.NewClient().Auth(APIKey, APISecret)
	return client
}

func (bfc *BitfinexClient) GetLendBook(currency string, limitBids, limitAsks int) bitfinex.Lendbook {
	book, err := bfc.Client.Lendbook.Get(currency, limitBids, limitAsks)
	if err != nil {
		fmt.Println("GetLendBook:", err)
		return bitfinex.Lendbook{}
	}
	return book
}

func (bfc *BitfinexClient) CreateLend(currency string, amount, rate float64, period int) bitfinex.MarginOffer {
	offer, err := bfc.Client.MarginFunding.NewLend(currency, amount, rate, period)
	if err != nil {
		fmt.Println("CreateLend:", err)
		return bitfinex.MarginOffer{}
	}
	return offer
}

func (bfc *BitfinexClient) GetFundingBalance() float64 {
	balance, err := bfc.Client.Balances.All()
	if err != nil {
		fmt.Println("GetFundingBalance:", err)
		return 0.0
	}
	for _, b := range balance {
		if b.Type == "deposit" && b.Currency == "usd" {
			available, _ := strconv.ParseFloat(b.Available, 64)
			s := fmt.Sprintf("%.3f", available)
			funding, _ := strconv.ParseFloat(s[:len(s)-1], 64)
			return funding
		}
	}
	return 0.0
}

func (bfc *BitfinexClient) GetAllFundingOffers() []bitfinex.ActiveOffer {
	offers, err := bfc.Client.MarginFunding.Offers()
	if err != nil {
		fmt.Println("GetAllFundingOffers:", err)
		return []bitfinex.ActiveOffer{}
	}
	return offers
}

func (bfc *BitfinexClient) CancelOffer(orderID int64) bitfinex.MarginOffer {
	v, err := bfc.Client.MarginFunding.Cancel(orderID)
	if err != nil {
		fmt.Println("CancelOffer:", err)
		return bitfinex.MarginOffer{}
	}
	return v
}
