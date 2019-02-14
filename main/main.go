package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func main() {

	//go server()
	go client()
	var a string
	fmt.Scanln(&a)

}

func client() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msg, err := ch.Consume(q.Name, "", true, false, false, true, nil)
	failOnErr(err, "Failed to consume messages")
	for mess := range msg {
		log.Printf("recieveved %v", mess.Body)
	}
}

func server() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msg := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hey there Anubhav"),
	}

	for {
		ch.Publish("", q.Name, false, false, msg)
	}

}

func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	conn, err := amqp.Dial("amqp://guest@localhost:5672")
	failOnErr(err, "Failed to connect to the localhost")
	ch, err := conn.Channel()
	failOnErr(err, "Failed to open a channel")
	q, err := ch.QueueDeclare("sensor", false, false, false, false, nil)
	failOnErr(err, "Failed to get the Queue")
	return conn, ch, &q

}

func failOnErr(err error, msg string) {
	if err != nil {
		log.Fatalf("%s : %s", err, msg)
		panic(fmt.Sprintf("%s : %s", err, msg))
	}
}
