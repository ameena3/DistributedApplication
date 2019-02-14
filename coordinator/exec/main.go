package main

import (
	"DistributedApplication/coordinator"
	"fmt"
)

func main() {
	ql := coordinator.NewQueueListner()
	go ql.ListenForNewSources()
	var a string
	fmt.Scanln(&a)
}
