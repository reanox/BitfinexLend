package service

import (
	"context"
	"fmt"
	"github.com/reanox/BitfinexLend/service/apiserver"
	"github.com/reanox/BitfinexLend/service/dbservice"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func RunAllService() {

	log.Println("Starting DB Service...")
	dbservice.Init()

	log.Println("Starting API Service...")
	server := apiserver.New()

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
