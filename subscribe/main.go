package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	NatsURL = "nats://localhost:14222"
)

func main() {
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	nc, err := nats.Connect(NatsURL,
		nats.UserInfo("nats", "root"),
		nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(10),
		nats.ReconnectWait(time.Minute),
		nats.PingInterval(time.Minute),
	)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()

	js, err := nc.JetStream()
	if err != nil {
		log.Fatalln(err)
	}

	sub, err := js.QueueSubscribe("TEST.*.*", "TEST", func(msg *nats.Msg) {
		fmt.Println("============================== Wildcard ==============================")
		fmt.Println("Subject", msg.Subject)
		fmt.Println("Reply", msg.Reply)
		fmt.Println("Header:", msg.Header)
		fmt.Println("Sub.Subject:", msg.Sub.Subject)
		fmt.Println("msg.Sub.Queue:", msg.Sub.Queue)
		fmt.Println("msg.Sub.Type:", msg.Sub.Type())
		fmt.Println("Data:", string(msg.Data))
		msg.Ack()
	}, nats.Durable("TEST"), nats.ManualAck(), nats.DeliverLast())
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Subject:", sub.Subject)

	select {
	case sig := <-stop:
		fmt.Printf("Service shutdown because %+v \n", sig)
		//sub.Drain()
		//sub.Unsubscribe()
		os.Exit(0)
	}
}
