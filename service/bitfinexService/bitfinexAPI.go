package bitfinexService

import (
	"fmt"
	"github.com/bitfinexcom/bitfinex-api-go/v1"
)

func NewClient(APIKey string, APISecret string) *bitfinex.Client {
	client := bitfinex.NewClient().Auth(APIKey, APISecret)
	info, err := client.Account.Info()

	if err != nil {
		fmt.Println(err)
		return nil
	} else {
		fmt.Println(info)
		return client
	}
}
