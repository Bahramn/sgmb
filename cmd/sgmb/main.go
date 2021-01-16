package main

import (
	"fmt"
	"github.com/Bahramn/sgmb/config"
	"github.com/Bahramn/sgmb/internal/app"
	"github.com/Bahramn/sgmb/internal/logger"
	"log"
	"time"
)

func main() {
	conf, err := config.Load()
	if err != nil {
		log.Fatalf("unable to load config struct, %v", err)
	}
	logger.InitLogger(conf.Log)

	s := app.NewServer()

	go s.Run()

	go s.ServeTcp(conf.Server.TCP)

	go s.ServeUdp(conf.Server.UDP)

	go s.ServeHttp(conf.Server.HTTP)

	for t := range time.Tick(30 * time.Second) {
		s.CheckClientsByLastPingAt()
		fmt.Printf("%v  active connections : %d \n", t.Format("2006-01-02 3:4:5 pm"), s.NumberOfClients())
	}

}
