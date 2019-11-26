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
		qos      = flag.Int("qos", 2, "QoS for published messages")
		size     = flag.Int("size", 100, "Size of the messages payload (bytes)")
		//count    = flag.Int("count", 100, "Number of messages to send per client")
		clients  = flag.Int("clients", 22, "Number of clients to start")
		topicNum = flag.Int("topicNum", 7, "Number of different topics")
	)

func main() {
	fmt.Println("New Program")
	
	flag.Parse()
	if *clients < 1 {
		log.Fatalf("Invalid arguments: number of clients should be > 1, given: %v", *clients)
	}

	if *topicNum > *clients {
		log.Fatalf("Invalid arguments: number of topics should be <= clients, given: %v", *topicNum)
	}

	clients_per_topic := *clients / *topicNum 
	log.Printf("Clients per topic %d", clients_per_topic)
	//start := time.Now()
	for i := 0; i < *clients; i++ {
		log.Printf("Topic name %d", i / clients_per_topic)
		c := &Client{
			ClientId:         i,
			BrokerURL:  *broker,
			MsgTopic:   *topic,
			MsgSize:    *size,
			MsgQoS:     byte(*qos),
			TopicNum:   (i / clients_per_topic),
		}
		go c.Run()
		log.Printf("Starting client %d", i)

	}

	time.Sleep(2 * time.Second)


}