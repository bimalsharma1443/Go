package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"encoding/json"
)


// Users is
type Users struct {
	Id        int `gorm:"primary_key, AUTO_INCREMENT"`
	Name      string
	Email   string `gorm:"column:email"`
	Phone string `gorm:"column:phone"`
}

// Attendance is
type Attendance struct {
	Id      int `gorm:"primary_key, AUTO_INCREMENT"`
	StudentId    int
	Users Users `gorm:"foreignkey:ID;association_foreignkey:StudentId"`
}

// Data is
type Data struct {
	Id      int `gorm:"primary_key, AUTO_INCREMENT"`
	AttId    int
	Attendance Attendance `gorm:"foreignkey:ID;association_foreignkey:AttId"`
}

func main() {

	db, err := gorm.Open("mysql", "root:@/godb")

	// checking a err
	if err != nil {
		panic("failed to connect database")
	}

	data := []Data{}

	db.LogMode(true)

	db.Preload("Attendance").Preload("Attendance.Users").Find(&data)

	// db close
	defer db.Close()

	value,_ := json.Marshal(data)

	fmt.Println(string(value))
	

}
