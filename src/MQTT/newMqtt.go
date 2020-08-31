package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Data struct {
	DeviceID    string
	DeviceName  string
	ChannelName string
	Value       int
	Tx          int64
}

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
	client := connect("sub", uri)
	fmt.Println("start listning")
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}

func main() {
	uri := "tcp://159.65.154.234:1883"
	topic := "testMqtt"

	go listen(uri, topic)

	client := connect("iot_client", uri)
	timer := time.NewTicker(1 * time.Second)
	value := 0
	for range timer.C {
		value = value + rand.Intn(10)
		str, _ := json.Marshal(&Data{
			DeviceID:    "M1",
			DeviceName:  "er4567",
			ChannelName: "ACC_1",
			Value:       value,
			Tx:          time.Now().Unix(),
		})
		client.Publish(topic, 0, false, str)

	}
}
