package bitfinexService

import (
	"log"
	"strconv"
	"time"

	"github.com/bitfinexcom/bitfinex-api-go/v1"
	"github.com/go-pg/pg"
	"github.com/reanox/BitfinexLend/db/models"
)

type Service struct {
	bftClients []*BitfinexClient
	db         *pg.DB
}

func New(db *pg.DB) (s *Service) {
	return &Service{
		bftClients: []*BitfinexClient{},
		db:         db,
	}
}

func (s *Service) Start() {
	// Auto lending
	go func() {
		for {
			for _, c := range s.bftClients {
				lendbook := c.GetLendBook("usd", 50, 50)
				var lendRate float64
				if len(lendbook.Bids) == 0 {
					continue
				}
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
				t := time.NewTicker(time.Minute * time.Duration(1))
				<-t.C
			}
		}
	}()

	// Cancel offer
	go func() {
		for {
			for _, c := range s.bftClients {
				timestamp := time.Now().Unix()
				offers := c.GetAllFundingOffers()
				for _, offer := range offers {
					offerT, _ := strconv.ParseFloat(offer.Timestamp, 64)
					log.Println(timestamp, offerT)
					if timestamp-int64(offerT) >= offerRemoveTime {
						result := c.CancelOffer(offer.ID)
						log.Printf("Remove offer %v", result)
						time.Sleep(time.Second * time.Duration(30))
					}
				}
			}
			t := time.NewTicker(time.Minute * time.Duration(5))
			<-t.C
		}
	}()
}

func (s *Service) CreateNewBFClient(APIKey string, APISecret string) *BitfinexClient {
	c := &BitfinexClient{
		Client: NewClient(APIKey, APISecret),
		User: &models.User{
			APIKey:        APIKey,
			APISecret:     APISecret,
			MinRate:       make(map[string]float64),
			ReserveFunds:  make(map[string]float64),
			MaxLendAmount: make(map[string]float64),
			MinDuration:   make(map[string]float64),
			MaxDuration:   make(map[string]float64),
		},
	}
	s.addClient(c)
	log.Printf("CreateNewBFClient: APIKey=%s", APIKey)
	return c
}

func (s *Service) removeClient(b *BitfinexClient) {
	index := 0
	for i, _b := range s.bftClients {
		if _b == b {
			index = i
			break
		}
	}
	s.bftClients = append(s.bftClients[:index], s.bftClients[index+1:]...)
}

func (s *Service) removeClientByKey(key string) {
	index := 0
	for i, _b := range s.bftClients {
		if _b.User.APIKey == key {
			index = i
			break
		}
	}
	s.bftClients = append(s.bftClients[:index], s.bftClients[index+1:]...)
}

func (s *Service) addClient(bfc *BitfinexClient) {
	s.bftClients = append(s.bftClients, bfc)
}
