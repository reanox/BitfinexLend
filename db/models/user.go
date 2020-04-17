package models

type User struct {
	APIKey        string
	APISecret     string
	MinRate       map[string]float64
	ReserveFunds  map[string]float64
	MaxLendAmount map[string]float64
	MinDuration   map[string]float64
	MaxDuration   map[string]float64
}
