package bitfinexService

import (
	"log"
	"strconv"
	"time"

	"github.com/bitfinexcom/bitfinex-api-go/v1"
)

func Init() {
	go func() {
		for {
			for _, c := range BFClients {
				lendbook, _ := c.GetLendBook("usd", 50, 50)
				var lendRate float64
				lendBidRate, _ := strconv.ParseFloat(lendbook.Bids[0].Rate, 64)
				if lendBidRate >= AnnualizedRate30d {
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
						if totalAmount > 20000.0 {
							break
						}

					}
				}
				log.Println("Get lend rate...", lendRate)

				fundingBalance, _ := c.GetFundingBalance()
				log.Println("Get funding balance...", fundingBalance)

				if fundingBalance <= miniumLendNumber || lendRate < AnnualizedRateMin {
					log.Println("Balance or Rate is low then minium.")
				} else {
					var offer bitfinex.MarginOffer
					if lendRate >= AnnualizedRate30d {
						offer, _ = c.CreateLend("USD", fundingBalance, lendRate, 30)
					} else {
						offer, _ = c.CreateLend("USD", fundingBalance, lendRate, 2)
					}
					log.Println("Create new offer:", offer)
				}
				time.Sleep(10 * time.Second)
			}
		}
	}()
}
