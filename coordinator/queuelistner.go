package coordinator

import (
	"DistributedApplication/dto"
	qutils "DistributedApplication/utils"
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/streadway/amqp"
)

const url = "amqp://guest:guest@localhost:5672"

//Queuelistner ...
type Queuelistner struct {
	con     *amqp.Connection
	ch      *amqp.Channel
	sources map[string]<-chan amqp.Delivery
	ea      *EventAggregrator
}

//NewQueueListner ...
func NewQueueListner() *Queuelistner {
	ql := Queuelistner{
		sources: make(map[string]<-chan amqp.Delivery),
		ea:      NewEventAggregrator(),
	}
	ql.con, ql.ch = qutils.GetChannel(url)
	return &ql
}

//ListenForNewSources ...
func (ql *Queuelistner) ListenForNewSources() {
	q := qutils.GetQueue("", ql.ch)
	ql.ch.QueueBind(q.Name,
		"",
		"amq.fanout",
		false,
		nil)

	msgs, err := ql.ch.Consume(q.Name, "", true, false, false, false, nil)
	qutils.FailOnErr(err, "Failed to counsume the message")

	for msg := range msgs {
		sourceChan, _ := ql.ch.Consume(string(msg.Body), "", true, false, false, false, nil)
		if ql.sources[string(msg.Body)] == nil {
			ql.sources[string(msg.Body)] = sourceChan
			go ql.AddListner(sourceChan)
		}
	}

}

//AddListner ...
func (ql *Queuelistner) AddListner(msgs <-chan amqp.Delivery) {
	for msg := range msgs {
		r := bytes.NewReader(msg.Body)
		d := gob.NewDecoder(r)
		sd := new(dto.SensorMessage)
		d.Decode(sd)
		fmt.Printf("Recieved : %v\n", sd)

		ed := EventData{
			Name:      sd.Name,
			Value:     sd.Value,
			TimeStamp: sd.Timestamp,
		}

		ql.ea.PublishEvent("MessageRecieved_"+msg.RoutingKey, ed)
	}

}
