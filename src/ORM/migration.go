package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// creating a model
type DeviceLog struct {
	gorm.Model
	DeviceId        string `gorm:"type:varchar(100)"`
	DeviceChannelId string `gorm:"type:int"`
	Value           string `gorm:"type:text"`
	CreatedOn       int64  `gorm:"type:varchar(100)"`
}

type DeviceModelChannel struct {
	gorm.Model
	Name        string `gorm:"type:varchar(100)"`
	Description string `gorm:"type:text"`
	DeviceId    string `gorm:"type:varchar(100)"`
}

func main() {
	// connect to database
	db, err := gorm.Open("mysql", "root:@/iot")

	// checking a err
	if err != nil {
		panic("failed to connect database")
	}

	// db close
	defer db.Close()

	// Migrate the schema
	db.AutoMigrate(&DeviceLog{})
	db.AutoMigrate(&DeviceModelChannel{})

}
