package main

import (
	"net"
	"time"

	log "github.com/Sirupsen/logrus"
	marketdata "github.com/ilovelili/sumoproto/services/marketdata/shared"
	"github.com/ilovelili/sumoproto/services/marketdata/subscriber/subscriber"

	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/transport/nats"
)

func opts(o *micro.Options) {
	o.Server = server.NewServer(func(o *server.Options) {
		o.Name = "invast.sumo.srv.marketdatasubscriber"
	})
}

func main() {
	cmd.Init()
	log.Println("Starting up marketdata subscriber...")
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Println(err)
	}

	log.Println("Interfaces:")
	for _, add := range addrs {
		log.Println(add.Network()+":", add.String())
	}
	s := micro.NewService(opts)

	if err = s.Server().Subscribe(
		server.NewSubscriber(
			marketdata.MarketDataServiceName,
			new(subscriber.MarketData),
		),
	); err != nil {
		log.Fatal(err)
	}

	retry := time.NewTicker(1 * time.Second)
RetryLoop:
	for {
		select {
		case <-retry.C:
			if err = s.Options().Broker.Connect(); err != nil {
				log.Error(err)
			} else {
				retry.Stop()
				break RetryLoop
			}
		}
	}

	if err = s.Run(); err != nil {
		log.Error(err)
	}
}
