package main

import (
	"log"
	"net"
	"runtime"
	"time"

	"github.com/bluefoxcode/rome/services/Caesar/lib/boot"
	"github.com/bluefoxcode/rome/services/Caesar/subscriber"
	micro "github.com/micro/go-micro"
	"github.com/micro/go-micro/cmd"
	"github.com/micro/go-micro/server"
	_ "github.com/micro/go-plugins/broker/nats"
	_ "github.com/micro/go-plugins/registry/nats"
	_ "github.com/micro/go-plugins/transport/nats"
)

func init() {
	log.SetFlags(log.Lshortfile)
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func handle() {
	log.Println("Ping received")
}

func main() {
	info := boot.LoadConfig()
	boot.RegisterServices(info)
	cmd.Init()
	log.Println("Starting up Notekeeper...")
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
			"go.db.Race.post",
			new(subscriber.Race),
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
				log.Println(err)
			} else {
				retry.Stop()
				break RetryLoop
			}
		}
	}

	if err = s.Run(); err != nil {
		log.Println(err)
	}
}

func opts(o *micro.Options) {
	o.Server = server.NewServer(func(o *server.Options) {
		o.Name = "go.db.Race.post"
	})
}
