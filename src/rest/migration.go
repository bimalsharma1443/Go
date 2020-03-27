package main

import (
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

	// Migrate the schema
	db.AutoMigrate(&User{})

}
