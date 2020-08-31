package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
	// mq "udp/server/mqttcode"
)

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func main() {

	PORT := ":9001"

	s, err := net.ResolveUDPAddr("udp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}

	connection, err := net.ListenUDP("udp4", s)
	if err != nil {
		fmt.Println(err)
		return
	}

	defer connection.Close()
	buffer := make([]byte, 1024)
	rand.Seed(time.Now().Unix())

	for {
		n, addr, err := connection.ReadFromUDP(buffer)
		fmt.Print("-> ", string(buffer[0:n-1]))

		if strings.TrimSpace(string(buffer[0:n])) == "STOP" {
			fmt.Println("Exiting UDP server!")
			return
		}

		data := []byte(strconv.Itoa(random(1, 1001)))
		fmt.Printf("data: %s\n", string(data))

		var sd SensorData
		json.Unmarshal([]byte(string(buffer[0:n-1])), &sd)
		fmt.Println(sd.Deviceid)
		fmt.Println(sd.Channel)
		fmt.Println(sd.Productioncount)

		mq.Publishdata("facemaskkey/"+sd.Deviceid, string(buffer[0:n-1]))

		_, err = connection.WriteToUDP(data, addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

//SensorData struct
type SensorData struct {
	Deviceid        string `json:"device_id"`
	Channel         string `json:"channel"`
	Productioncount int    `json:"count"`
}
