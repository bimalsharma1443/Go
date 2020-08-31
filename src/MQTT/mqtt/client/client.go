package main

import (
	"fmt"
	"log"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func connect(clientId, uri string) mqtt.Client {
	fmt.Println("start connect")
	opts := createClientOptions(clientId, uri)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

func createClientOptions(clientId, uri string) *mqtt.ClientOptions {
	fmt.Println("start create client")
	opts := mqtt.NewClientOptions()
	opts.AddBroker(uri)
	opts.SetUsername("admin")
	password := "public"
	opts.SetPassword(password)
	opts.SetClientID(clientId)
	return opts
}

func listen(uri, topic string) {
	client := connect("iot_client", uri)
	fmt.Println("start listning")
	for {
		tt := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
			fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
		})
		fmt.Println(tt)
	}

}

func main() {
	uri := "tcp://broker.emqx.io:1883"
	topic := "iot"

	listen(uri, topic)

}
