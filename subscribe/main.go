package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nats-io/go-nats-streaming"
	"github.com/sirupsen/logrus"
)

const (
	NatsURL = "nats://nats:nats@localhost:4222"
)

func main() {
	stop := make(chan os.Signal)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)
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

	logrus.WithField("client_id", clientID).Infoln("Subscriber start")

	// Subscribe message with group
	sub, err := sc.QueueSubscribe("foo", "group-1", func(msg *stan.Msg) {
		logrus.WithFields(logrus.Fields{
			"timestamp":   msg.Timestamp,
			"redelivered": msg.Redelivered,
			"reply":       msg.Reply,
			"subject":     msg.Subject,
			"data":        string(msg.Data),
		}).Infoln("QueueSubscribe message from server")
		msg.Ack()
	}, stan.DurableName("test-123"), stan.SetManualAckMode())
	if err != nil {
		logrus.WithError(err).Fatalln("Cannot subscribe")
	}
	logrus.Infoln(sub.IsValid())

	select {
	case sig := <-stop:
		sub.Unsubscribe()
		logrus.Warningln(fmt.Sprintf("Service shutdown because %+v", sig))
		os.Exit(0)
	}
}
