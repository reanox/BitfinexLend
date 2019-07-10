package db

type User struct {
	APIKey       string
	APISecret    string
	MinRate      map[string]float64
	ReserveQuota map[string]float64
}
