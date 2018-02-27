package main

import (
	"net"
	"time"

	log "github.com/Sirupsen/logrus"

	"github.com/ilovelili/sumoproto/services/tickrecorder/subscriber"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/transport/nats"
)

func opts(o *micro.Options) {
	o.Server = server.NewServer(func(o *server.Options) {
		o.Name = "go.micro.srv.tickrecorder"
	})
}

func handle() {
	log.Println("Tick received")
}

func main() {
	cmd.Init()
	log.Println("Starting up Tick Recorder...")
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
			"go.micro.srv.TickRecorder.Tick",
			new(subscriber.Tick),
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
