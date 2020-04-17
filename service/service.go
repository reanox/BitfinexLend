package service

import (
	"context"
	"fmt"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	serviceconfig "github.com/reanox/BitfinexLend/config"
	"github.com/reanox/BitfinexLend/db"
	"github.com/reanox/BitfinexLend/service/apiserver"
	"github.com/reanox/BitfinexLend/service/bitfinexService"
)

func RunAllService() {

	log.Println("Starting DB Service...")
	db := db.NewDB(serviceconfig.DBUser, serviceconfig.DBPassword, serviceconfig.DBName)

	log.Println("Starting Bitfinex Lend Service...")
	bftServiceInstance := bitfinexService.New(db)
	bftServiceInstance.Start()

	log.Println("Starting API Service...")
	server := apiserver.New(bftServiceInstance)

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			e := fmt.Errorf("Listen http failed. err=[%v]", err)
			fmt.Println(e)
			return
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		log.Println("Server Shutdown:", err)
	}
}
