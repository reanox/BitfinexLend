package apiserver

import (
	"github.com/reanox/BitfinexLend/router"
	"net/http"
)

// New a server
func New() *http.Server {
	server := &http.Server{
		Addr:    ":8080",
		Handler: router.New(),
	}
	return server
}
