package main

import (
	"fmt"
	"flag"
	"log"
	"time"
)

var (
		//resultChan = make(chan Result)

		broker   = flag.String("broker", "tcp://localhost:1883", "MQTT broker endpoint as scheme://host:port")
		topic    = flag.String("topic", "test", "MQTT topic for outgoing messages")
		qos      = flag.Int("qos", 1, "QoS for published messages")
		size     = flag.Int("size", 100, "Size of the messages payload (bytes)")
		//count    = flag.Int("count", 100, "Number of messages to send per client")
		clients  = flag.Int("clients", 10, "Number of clients to start")
	)

func main() {
	fmt.Println("New Program")
	
	flag.Parse()
	if *clients < 1 {
		log.Fatalf("Invalid arguments: number of clients should be > 1, given: %v", *clients)
	}

	//start := time.Now()
	for i := 0; i < *clients; i++ {
		c := &Client{
			ClientId:         i,
			BrokerURL:  *broker,
			MsgTopic:   *topic,
			MsgSize:    *size,
			MsgQoS:     byte(*qos),
		}
		go c.Run()
		log.Println("Starting client ", i)

	}

	time.Sleep(1 * time.Second)


}