package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/nats-io/go-nats-streaming"
	"github.com/sirupsen/logrus"
)

const (
	NatsURL = "nats://nats:nats@localhost:4222"
)

func main() {
	clientID := time.Now().Unix()

	sc, err := stan.Connect("dynastymasra-cluster", fmt.Sprintf("%v", clientID),
		stan.NatsURL(NatsURL),
		stan.ConnectWait(time.Minute),
		stan.Pings(60, 5),
		stan.SetConnectionLostHandler(func(conn stan.Conn, err error) {
			logrus.Warningln("NATS connection is lost")
			if err != nil {
				logrus.WithError(err).Fatalln("Error connection lost")
			}
		}))
	if err != nil {
		logrus.WithError(err).Fatalln("Cannot connect to nats cluster")
	}
	defer sc.Close()

	logrus.WithField("client_id", clientID).Infoln("Publisher start")

	r := rand.Intn(10)
	for i := 0; i < 10; i++ {
		// Simple Synchronous Publisher
		if err := sc.Publish("foo", []byte(fmt.Sprintf("Hello World Synchronous %d", i))); err != nil {
			logrus.WithError(err).Errorln("Failed publish message")
		}

		time.Sleep(time.Duration(r) * time.Second)
	}

	for i := 0; i < 10; i++ {
		// Simple Asynchronous Publisher
		nuid, err := sc.PublishAsync("foo", []byte(fmt.Sprintf("Hello World Asynchronous %d", i)), func(nuid string, err error) {
			if err != nil {
				logrus.WithError(err).Errorln("Failed publish message NUID " + nuid)
			} else {
				logrus.Infoln("Received ack for message nuid " + nuid)
			}
		})
		if err != nil {
			logrus.WithError(err).Errorln("Failed publish message")
		}
		logrus.Infoln("Asynchronous NUID = " + nuid)

		time.Sleep(time.Duration(r) * time.Second)
	}
}
