package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"time"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

// IOTData set info
type IOTData struct {
	DeviceID    string
	DeviceName  string
	ChannelName string
	Value       int
	Tx          int64
}

func main() {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker("tcp://159.65.154.234:1883")
	opts.SetClientID("Device-pub")

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)

	//we are going to try connecting for max 10 times to the server if the connection fails.
	for i := 0; i < 10; i++ {
		if token := c.Connect(); token.Wait() && token.Error() == nil {
			break
		} else {
			fmt.Println(token.Error())
			time.Sleep(1 * time.Second)
		}
	}

	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	//same thing needs to go here as well.
	if token := c.Subscribe("some_topic", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// this is the shell where we will take input from the user and publish the message on the topic until user enters `bye`.
	value := 0
	for {
		value = value + rand.Intn(10)
		// if there is a message, publish it.
		message, _ := json.Marshal(&IOTData{
			DeviceID:    "M1",
			DeviceName:  "er4567",
			ChannelName: "ACC_1",
			Value:       value,
			Tx:          time.Now().Unix(),
		})

		token := c.Publish("some_topic", 0, false, message)
		fmt.Println("se :- " + string(message))
		token.Wait()
		time.Sleep(time.Second * 2)
	}

	//unsubscribe from /go-mqtt/sample
	if token := c.Unsubscribe("some_topic"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)

}
