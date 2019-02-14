package coordinator

import (
	"time"
)

//EventAggregrator ...
type EventAggregrator struct {
	listners map[string][]func(EventData)
}

//EventData ...
type EventData struct {
	Name      string
	Value     float64
	TimeStamp time.Time
}

//NewEventAggregrator ...
func NewEventAggregrator() *EventAggregrator {
	ea := EventAggregrator{
		listners: make(map[string][]func(EventData)),
	}
	return &ea
}

//AddListner ...
func (ea *EventAggregrator) AddListner(name string, f func(EventData)) {
	ea.listners[name] = append(ea.listners[name], f)
}

//PublishEvent ...
func (ea *EventAggregrator) PublishEvent(name string, ed EventData) {
	if ea.listners[name] != nil {
		for _, listner := range ea.listners[name] {
			listner(ed)
		}
	}
}
