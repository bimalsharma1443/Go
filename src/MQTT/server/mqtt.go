package main

import (
	"fmt"
	"strconv"

	//import the Paho Go MQTT library
	"os"
	"time"

	"github.com/go-redis/redis"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var flag bool = false

//var wcount int = 0

//define a function for the default message handler
var f MQTT.MessageHandler = func(client MQTT.Client, msg MQTT.Message) {
	topic := msg.Topic()
	payload := msg.Payload()

	clientRedis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	value := time.Now().Unix()
	key := topic + "_" + strconv.FormatInt(value, 10)

	err := clientRedis.Set(key, string(payload), 0).Err()
	fmt.Println(err)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("re :- " + string(payload))

}

func main() {
	//create a ClientOptions struct setting the broker address, clientid, turn
	//off trace output and set the default message handler
	opts := MQTT.NewClientOptions().AddBroker("tcp://159.65.154.234:1883")
	opts.SetClientID("Device-sub")
	opts.SetDefaultPublishHandler(f)

	//create and start a client using the above ClientOptions
	c := MQTT.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	//subscribe to the topic /go-mqtt/sample and request messages to be delivered
	//at a maximum qos of zero, wait for the receipt to confirm the subscription
	if token := c.Subscribe("some_topic", 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for flag == false {
		time.Sleep(1 * time.Second)
		//fmt.Println("waiting: ", wcount)
		//wcount += 1
	}

	//unsubscribe from /go-mqtt/sample
	if token := c.Unsubscribe("some_topic"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	c.Disconnect(250)
}
