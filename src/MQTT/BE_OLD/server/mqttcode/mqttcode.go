package mqttcode

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

/*var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	fmt.Printf("TOPIC: %s\n", msg.Topic())
	fmt.Printf("MSG: %s\n", msg.Payload())
}
*/
//Publishdata publish
func Publishdata(topic string, data string) {

	fmt.Println(topic)
	//mqtt.DEBUG = log.New(os.Stdout, "", 0)
	//mqtt.ERROR = log.New(os.Stdout, "", 0)
	//cf074d4055fb@172.17.0.2@broker.emqx.io:1883
	opts := mqtt.NewClientOptions().AddBroker("tcp://broker.emqx.io:1883").SetClientID("emqx_test_client")
	//opts := mqtt.NewClientOptions().AddBroker("mqtt://cf074d4055fb@172.17.0.2@broker.emqx.io:1883").SetClientID("emqx_test_client")

	opts.SetKeepAlive(60 * time.Second)
	// Set the message callback handler
	//opts.SetDefaultPublishHandler(f)
	//opts.SetPingTimeout(1 * time.Second)
	opts.SetUsername("mqttuser")
	opts.SetPassword("123456")
	//opts.SetProtocolVersion(3)

	//tlsConfig := &tls.Config{InsecureSkipVerify: false, ClientAuth: tls.NoClientCert}
	//opts.SetTLSConfig(tlsConfig)

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		fmt.Println("err", token.Error())
		panic(token.Error())
	}

	// Publish a message
	token := c.Publish(topic, 0, true, data)
	token.Wait()

	//fmt.Println(token)

	time.Sleep(6 * time.Second)
	fmt.Println("Published data")

	// Cancel subscription
	/*if token := c.Unsubscribe("facemaskkey/Test"); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}*/

	c.Disconnect(250)
	time.Sleep(1 * time.Second)
	//Subscribedata(topic, data)

}

//Subscribedata for subscribing
/*func Subscribedata(topic string, data string) {
	mqtt.DEBUG = log.New(os.Stdout, "", 0)
	mqtt.ERROR = log.New(os.Stdout, "", 0)
	opts := mqtt.NewClientOptions().AddBroker("tcp://broker.emqx.io:1883").SetClientID("emqx_test_client")

	opts.SetKeepAlive(60 * time.Second)
	// Set the message callback handler
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)
	opts.SetUsername("mqttuser")
	opts.SetPassword("123456")

	c := mqtt.NewClient(opts)
	if token := c.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	// Subscribe to a topic
	if token := c.Subscribe(topic, 0, nil); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	var sd db.Tele_data
	sd.Ts = time.Now().Unix()
	json.Unmarshal([]byte(string(data)), &sd)
	go db.Savedata(sd)
	//fmt.Println("subscribed the data")

	// Unscribe
	if token := c.Unsubscribe(topic); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	// Disconnect
	c.Disconnect(250)
	time.Sleep(1 * time.Second)
}*/
