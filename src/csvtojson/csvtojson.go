package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"
)

type stage struct {
	Stage          int    `json:"stage"`
	Period         int    `json:"period"`
	StartTimestamp int64  `json:"startTimestamp"`
	ActiveCode     string `json:"activeCode"`
}

type dongleFlashData struct {
	RemoteControl int     `json:"remoteControl"`
	Maxtries      int     `json:"maxtries"`
	Tries         int     `json:"tries"`
	Current       int     `json:"current"`
	Licenses      []stage `json:"licenses"`
}

type dongleData struct {
	DongleID         string `json:"dongleID"`
	AdminStageNumber int    `json:"adminStageNumber"`
	AdminStage       int    `json:"adminStage"`
	RemoteAPI        int    `json:"remoteAPI"`
	Batch            string `json:"batch"`
	DongleFlashData  string `json:"dongleFlashData"`
}

type data struct {
	DongleData []dongleData `json:"dongleData"`
}

func main() {
	csvfile, err := os.Open("20200630_all.csv")
	if err != nil {
		log.Fatalln("Couldn't open the csv file", err)
	}

	r := csv.NewReader(csvfile)

	dongleDataSlice := []dongleData{}

	count := 0

	for {
		// Read each record from csv

		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if count == 0 {
			count++
			continue
		}

		index := 3
		stageSlice := []stage{}
		for i := 0; i < 5; i++ {
			if record[index] == "" {
				break
			}

			stage := stage{
				Stage:          i + 1,
				Period:         0,
				StartTimestamp: toTimeStamp(record[index]),
				ActiveCode:     getToken(record[index+1]),
			}
			index += 2
			stageSlice = append(stageSlice, stage)
		}

		dongleFlashData := dongleFlashData{
			RemoteControl: 0,
			Maxtries:      10,
			Tries:         0,
			Current:       1,
			Licenses:      stageSlice,
		}

		dongleFlashDataString, _ := json.Marshal(dongleFlashData)

		remoteAPIint := 0
		if record[2] != "no" {
			remoteAPIint = 1
		}

		dongleData := dongleData{
			DongleID:         record[0],
			AdminStageNumber: 1,
			AdminStage:       1,
			RemoteAPI:        remoteAPIint,
			Batch:            record[1],
			DongleFlashData:  string(dongleFlashDataString),
		}

		dongleDataSlice = append(dongleDataSlice, dongleData)

	} // end for

	data := data{
		DongleData: dongleDataSlice,
	}

	// jsonData, err := json.Marshal(data)
	jsonData, err := json.MarshalIndent(data, "", " ")

	if err != nil {
		fmt.Println(err)
		return
	}

	// d1 := []byte(string(jsonData))
	err = ioutil.WriteFile("temp.json", jsonData, 0644)
	fmt.Println(err)

}

func toTimeStamp(date string) int64 {
	layout := "2006-01-02"
	t, _ := time.Parse(layout, date)
	return t.Unix()
}

func getToken(data string) string {
	return strings.ReplaceAll(data, "-", "")
}
