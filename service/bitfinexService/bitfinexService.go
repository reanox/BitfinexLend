package bitfinexService

import (
	"log"
	"strconv"
	"time"

	"github.com/bitfinexcom/bitfinex-api-go/v1"
)

func Start() {
	// Auto lending
	go func() {
		for {
			for _, c := range BFClients {
				lendbook := c.GetLendBook("usd", 50, 50)
				var lendRate float64
				lendBidRate, _ := strconv.ParseFloat(lendbook.Bids[0].Rate, 64)
				if lendBidRate >= annualizedRate30d {
					lendRate = lendBidRate
				}
				if lendRate == 0.0 {
					var totalAmount float64
					for _, ask := range lendbook.Asks {
						amount, _ := strconv.ParseFloat(ask.Amount, 64)
						totalAmount += amount
						if totalAmount > 500000.0 {
							break
						}
						lendRate, _ = strconv.ParseFloat(ask.Rate, 64)
						if totalAmount > 10000.0 {
							break
						}

					}
				}
				log.Println("Get lend rate...", lendRate)

				fundingBalance := c.GetFundingBalance()
				log.Println("Get funding balance...", fundingBalance)

				if fundingBalance <= miniumLendNumber || lendRate < annualizedRateMin {
					log.Println("Balance or Rate is low then minium.")
				} else {
					var offer bitfinex.MarginOffer
					if lendRate >= annualizedRate30d {
						offer = c.CreateLend("USD", fundingBalance, lendRate, 30)
					} else {
						offer = c.CreateLend("USD", fundingBalance, lendRate, 2)
					}
					log.Println("Create new offer:", offer)
				}
				time.Sleep(10 * time.Second)
			}
		}
	}()

	// Cancel offer
	go func() {
		for {
			for _, c := range BFClients {
				timestamp := time.Now().Unix()
				offers := c.GetAllFundingOffers()
				for _, offer := range offers {
					offerT, _ := strconv.ParseInt(offer.Timestamp, 10, 64)
					if timestamp-offerT >= offerRemoveTime {
						result := c.CancelOffer(offer.ID)
						log.Printf("Remove offer %v", result)
						time.Sleep(10 * time.Second)
					}
				}
			}
			time.Sleep(60 * time.Second)
		}
	}()
}
