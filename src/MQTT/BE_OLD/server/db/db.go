package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"

	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Savedata to save data in db
func Savedata(td Tele_data) {
	var db *gorm.DB
	// Set the file name of the configurations file
	viper.SetConfigName("config")

	// Set the path to look for the configurations file
	//viper.AddConfigPath(".")

	// Enable VIPER to read Environment Variables
	viper.AutomaticEnv()

	viper.SetConfigType("yml")

	viper.SetConfigFile("./config.yml")

	if err := viper.ReadInConfig(); err != nil {
		//level.Error(logger).Log("Error reading config file, %s", err)
		fmt.Println(err)
	}

	//url := "host=" + viper.GetString("database.host") + " port=" + viper.GetString("database.port") + " user=" + viper.GetString("database.user") + " dbname=" + viper.GetString("database.dbname") + " password=" + viper.GetString("database.password") + " sslmode=disable"

	url := "host=mighty-iot.xoxiot.com port=5432 user=postgres dbname=mighty_iot password=mighty123456 sslmode=disable"
	fmt.Println(url)
	db, dberr := gorm.Open("postgres", url)
	defer db.Close()
	if dberr != nil {
		fmt.Println(dberr)
		os.Exit(-1)
	}

	db.LogMode(true)
	//var dt Tele_data
	//dt.Device_id = "TestUDP"
	//dt.Channel = "Testchannel"
	//dt.Ts = time.Now().Unix()
	//dt.Dbl_v = 214523
	//db.Raw("SELECT (MAX(dbl_v)-MIN(dbl_v)) as count,MAX(ts) as timestamp,device_id as deviceid FROM tele_data where channel=? and ts>=? Group by device_id", "accYield", "1594578600000").Scan(&dt)
	db.Debug().Create(&td)
}

//Tele_data struct for details
type Tele_data struct {
	Device_id string  `gorm:"device_id" json:"device_id"`
	Channel   string  `gorm:"channel" json:"channel"`
	Ts        int64   `gorm:"ts" json:"ts"`
	Bool_v    bool    `gorm:"bool_v" json:"bool_v"`
	Str_v     string  `gorm:"str_v" json:"str_v"`
	Long_v    int64   `gorm:"long_v" json:"long_v"`
	Dbl_v     float64 `gorm:"dbl_v" json:"dbl_v"`
}

//{"device_id":"TestSErver", "channel":"Testserver","Dbl_v":12315}
