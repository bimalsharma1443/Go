package main

import (
	"encoding/json"
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Faculty is
type Faculty struct {
	Id      int `gorm:"primary_key, AUTO_INCREMENT"`
	Name    string
	Student []Student `gorm:"ForeignKey:FacultyID"`
}

// Student is
type Student struct {
	Id        int `gorm:"primary_key, AUTO_INCREMENT"`
	Name      string
	Address   string
	FacultyID int `gorm:"column:faculty_id"`
	Faculty   Faculty
}

func main() {
	fmt.Println("starting a code ----------------------------------")

	// connect to database
	db, err := gorm.Open("mysql", "root:@/godemo")

	// checking a err
	if err != nil {
		panic("failed to connect database")
	}

	// db close
	defer db.Close()

	//create user struct object
	faculties := Faculty{}
	db.LogMode(true)

	db.Preload("Student").Find(&faculties)

	fmt.Println(err)

	data, _ := json.Marshal(faculties)

	fmt.Println(string(data))

	fmt.Println("end of code --------------------------------------")

}
