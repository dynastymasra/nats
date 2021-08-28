package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/nats-io/nats.go"
)

const (
	NatsURL = "nats://localhost:14222"
)

func main() {
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

	as, err := js.AddStream(&nats.StreamConfig{
		Name:        "test",
		Description: "Event stream from test",
		Subjects:    []string{"TEST.*.*"},
	})
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("Name:", as.Config.Name)
	fmt.Println("Description:", as.Config.Description)
	fmt.Println("Subjects:", as.Config.Subjects)

	r := rand.Intn(10)
	for i := 0; i < 10; i++ {
		// Simple Synchronous Publisher
		fmt.Println("============================== PublishMsg ==============================")
		msg := nats.NewMsg("TEST.sync.test")
		msg.Data = []byte(fmt.Sprintf("Hello World Synchronous %d", i))
		msg.Header.Add("request_id", fmt.Sprintf("%d", i))
		pub, err := js.PublishMsg(msg, nats.Context(context.TODO()))
		if err != nil {
			fmt.Printf("PublishMsg: %d - %v \n", i, err)
			continue
		}

		fmt.Println("Domain:", pub.Domain)
		fmt.Println("Stream:", pub.Stream)
		fmt.Println("Duplicate:", pub.Duplicate)
		fmt.Println("Sequence:", pub.Sequence)

		time.Sleep(time.Duration(r) * time.Second)
	}

	for i := 0; i < 10; i++ {
		// Simple Asynchronous Publisher
		fmt.Println("============================== PublishMsgAsync ==============================")
		msg := nats.NewMsg("TEST.async.test")
		msg.Data = []byte(fmt.Sprintf("Hello World Asyncronous %d", i))
		msg.Header.Add("request_id", fmt.Sprintf("%d", i))
		pub, err := js.PublishMsgAsync(msg)
		if err != nil {
			fmt.Printf("PublishMsgAsync: %d - %v \n", i, err)
			continue
		}

		fmt.Println("Message:", pub.Msg())
	}
}
