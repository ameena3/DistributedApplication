package qutils

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

// SensorListQueue ...
const SensorListQueue = "SensorList"

// GetChannel ...
func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	FailOnErr(err, "Failed to create the connection.")
	ch, err := conn.Channel()
	FailOnErr(err, "Failed to create the Channel")
	return conn, ch
}

//GetQueue ...
func GetQueue(name string, ch *amqp.Channel) *amqp.Queue {
	q, err := ch.QueueDeclare(name, false, false, false, false, nil)
	FailOnErr(err, "Failed to create the Queue")
	return &q
}

//FailOnErr ...
func FailOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", err, msg)
		panic(fmt.Sprintf("%s : %s", err, msg))
	}
}
