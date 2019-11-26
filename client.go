package main

import (
	//"fmt"
	"log"
	//"time"
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
}

func (c *Client) Run() {

	subscriberOptions := mqtt.NewClientOptions().AddBroker(c.BrokerURL)
	subscriber := mqtt.NewClient(subscriberOptions)

	log.Printf("[%d] connecting subscriber\n", c.ClientId)

	if token := subscriber.Connect(); token.WaitTimeout(60) && token.Error() != nil {
		
		log.Printf("Connecting error\n", token)

		return
	}
}