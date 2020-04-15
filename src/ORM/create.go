package main

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// creating a model
type User struct {
	gorm.Model
	Name  string
	Age   string `gorm:"type:varchar(10)"`
	Email string `gorm:"type:varchar(100)"`
}

func main() {
	// connect to database
	db, err := gorm.Open("mysql", "root:@/godb?charset=utf8&parseTime=True&loc=Local")

	// checking a err
	if err != nil {
		panic("failed to connect database")
	}

	// db close
	defer db.Close()

	// value
	userData := User{
		Name:  "Shreya",
		Age:   "27",
		Email: "demo@demo.com"}

	// create row
	db.Create(&userData)

	// new Record inserted
	record := db.NewRecord(userData)

	fmt.Println("Record : ", record)

}
