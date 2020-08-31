package main

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DeviceModelChannel a model
type DeviceModelChannel struct {
	ID          int
	Name        string `gorm:"name"`
	Description string `gorm:"description"`
	DeviceID    string `gorm:"device_id"`
}

// DeviceLog is
type DeviceLog struct {
	DeviceID           string             `gorm:"device_id" json:"device_id"`
	DeviceChannelID    int                `gorm:"device_channel_id" json:"device_channel_id"`
	DeviceModelChannel DeviceModelChannel `gorm:"foreignkey:id;association_foreignkey:device_channel_id" json:"device_model_channel"`
	Value              int                `gorm:"value" json:"value"`
	CreatedOn          int64              `gorm:"created_on" json:"created_on"`
}

func main() {
	// connect to database
	db, err := gorm.Open("mysql", "root:@/iot")
	db.LogMode(true)
	// checking a err
	if err != nil {
		panic("failed to connect database")
	}

	// db close
	defer db.Close()

	deviceLog := DeviceLog{}

	DB :=
		db.
			Preload("DeviceModelChannel").
			Order("id desc").
			Limit("1").
			Find(&deviceLog)

	data, _ := json.Marshal(deviceLog)

	fmt.Println(deviceLog)
	fmt.Println(string(data))
	fmt.Println(DB)

}
