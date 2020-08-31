package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"time"
	"udp/server/db"

	//"log"
	"os"
	"os/signal"
	"syscall"

	MQTT "github.com/eclipse/paho.mqtt.golang"
)

//var wg sync.WaitGroup

func onMessageReceived(client MQTT.Client, message MQTT.Message) {
	//	wg.Add(1)
	//	go func() {
	//	defer wg.Done()
	//fmt.Printf("Received message on topic: %s\nMessage: %s\n", message.Topic(), message.Payload())
	fmt.Printf("%s\n", message.Payload())
	var sd db.Tele_data
	sd.Ts = time.Now().Unix()
	json.Unmarshal([]byte(message.Payload()), &sd)
	fmt.Println(sd)

	//go db.Savedata(sd)
	//}()

}

func main() {
	//MQTT.DEBUG = log.New(os.Stdout, "", 0)
	//MQTT.ERROR = log.New(os.Stdout, "", 0)
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	//hostname, _ := os.Hostname()

	server := flag.String("server", "tcp://broker.emqx.io:1883", "The full url of the MQTT server to connect to ex: tcp://127.0.0.1:1883")
	topic := flag.String("topic", "facemaskkey/Test", "Topic to subscribe to")

	qos := flag.Int("qos", 0, "The QoS to subscribe to messages at")
	//clientid := flag.String("clientid", hostname+strconv.Itoa(time.Now().Second()), "A clientid for the connection")

	flag.Parse()
	connOpts := MQTT.NewClientOptions().AddBroker("tcp://broker.emqx.io:1883").SetClientID("emqx_test_client")

	//connOpts := MQTT.NewClientOptions().AddBroker("mqtt://cf074d4055fb@172.17.0.2@broker.emqx.io:1883").SetClientID("emqx_test_client").SetCleanSession(true)
	connOpts.SetUsername("mqttuser")
	connOpts.SetPassword("123456")
	connOpts.SetPingTimeout(1 * time.Second)
	//connOpts := MQTT.NewClientOptions().AddBroker(*server).SetClientID(*clientid).SetCleanSession(true)

	//tlsConfig := &tls.Config{InsecureSkipVerify: true, ClientAuth: tls.NoClientCert}
	//connOpts.SetTLSConfig(tlsConfig)

	connOpts.OnConnect = func(c MQTT.Client) {
		fmt.Println("in onconect")
		if token := c.Subscribe(*topic, byte(*qos), onMessageReceived); token.Wait() && token.Error() != nil {
			fmt.Println("in err ", token.Error())
			panic(token.Error())
		}
	}

	client := MQTT.NewClient(connOpts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	} else {
		fmt.Printf("Connected to %s\n", *server)
	}

	<-c
}
