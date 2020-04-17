package types

// Response : Format return
type Response struct {
	Message   string `json:"message"`
	Status    int    `json:"status"`
	ErrorCode int    `json:"errorCode"`
}
