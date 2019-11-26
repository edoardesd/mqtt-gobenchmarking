package main

import (
	"fmt"
	"log"
	"time"
)

import (
	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	ClientId   int
	BrokerURL  string
	MsgTopic   string
	MsgSize    int
	MsgQoS     byte
	TopicNum   int
}

func (c *Client) Run() {

	subscriberOptions := mqtt.NewClientOptions().AddBroker(c.BrokerURL)
	subscriber := mqtt.NewClient(subscriberOptions)
	topicName := fmt.Sprintf("topic/%d", c.TopicNum)

	log.Printf("[%d] connecting subscriber\n", c.ClientId)

	if token := subscriber.Connect(); token.WaitTimeout(60* time.Second) && token.Error() != nil {

		return
	}

	log.Printf("[%d] subscribing to topic %s\n", c.ClientId, topicName)
	if token := subscriber.Subscribe(topicName, c.MsgQoS, nil); token.WaitTimeout(60* time.Second) && token.Error() != nil {
		log.Printf("Subscribing error\n", token)

		return
	}
}