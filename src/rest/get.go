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
	fmt.Println("starting a code ----------------------------------")

	// connect to database
	db, err := gorm.Open("mysql", "root:@/godb?charset=utf8&parseTime=True&loc=Local")

	// checking a err
	if err != nil {
		panic("failed to connect database")
	}

	// db close
	defer db.Close()

	//create user struct object
	userFirst := User{}
	// select * from users order by id LIMIT 1
	fmt.Println("-------------Find first user---------------------")
	db.First(&userFirst)
	fmt.Println(userFirst)
	fmt.Println("----------------------------------")

	fmt.Println("-------------take user---------------------")
	userTake := User{}
	db.Take(&userTake)
	fmt.Println(userTake)
	fmt.Println("----------------------------------")

	fmt.Println("-------------get last user---------------------")
	userLast := User{}
	db.Last(&userLast)
	fmt.Println(userLast)
	fmt.Println("----------------------------------")

	fmt.Println("-------------Find  user---------------------")
	userAll := User{}
	db.Find(&userAll)
	fmt.Println(userAll)
	fmt.Println("----------------------------------")

	fmt.Println("-------------Find  users---------------------")
	userFirstNumber := User{}
	db.First(&userFirstNumber, 5)
	fmt.Println(userFirstNumber)
	//// SELECT * FROM users WHERE id = 5;

	fmt.Println("------------------------find a user------------------------")
	userWhere := User{}
	db.Where("name = ?", "bimal").Find(&userWhere)
	fmt.Println(userWhere)

	fmt.Println("------------------------first a user------------------------")
	userFirstWhere := User{}
	db.Where("name = ?", "bimal").First(&userFirstWhere)
	fmt.Println(userFirstWhere)

	fmt.Println("------------------------In a user------------------------")
	userInWhere := []User{}
	db.Where("age IN(?)", []int{30, 27}).Find(&userInWhere)
	fmt.Println(userInWhere)

	fmt.Println("------------------------In a user------------------------")
	userBetweenWhere := []User{}
	db.Where("age between ? AND ?", 21, 30).Find(&userBetweenWhere)
	fmt.Println(userBetweenWhere)
	fmt.Println("----------------------------------")

	fmt.Println("------------------ struct and Map ------------------------")
	userWithObject := User{}
	db.Where(&User{Name: "bimal", Age: "27"}).Find(&userWithObject)
	fmt.Println(userWithObject)

	fmt.Println("-------------------------------- use of not ---------------------------")
	userNotData := User{}
	db.Not("name", "bimal").Find(&userNotData)
	fmt.Println(userNotData)
	fmt.Println("-----------------------------------------------------")

	fmt.Println("--------------------------------------- use of OR ---------------------")
	userOrData := []User{}
	db.Where("name = ?", "Bimal").Or("age = ?", 27).Find(&userOrData)
	fmt.Println(userOrData)
	fmt.Println("----------------------------------------------------------")

	fmt.Println("---------------------------- inline db -----------------------------")
	userfindData := User{}
	db.First(&userfindData, 5)
	fmt.Println(userfindData)
	fmt.Println("----------------------------------------------------------")
	userfindData = User{}
	db.Find(&userfindData, "name <> ? AND age > ?", "Bimal", 30)
	fmt.Println(userfindData)
	fmt.Println("----------------------------------------------------------")

	fmt.Println("--------------------------------------- first and initi ---------------------")
	user := User{}
	db.FirstOrInit(&user, User{Name: "demo"})
	fmt.Println(user)

	user = User{}
	db.Where(User{Name: "shiksha"}).Attrs(User{Age: "23"}).FirstOrInit(&user)
	fmt.Println(user)

	user = User{}
	db.Where(User{Name: "Bimal"}).Attrs(User{Age: "30"}).FirstOrInit(&user)
	fmt.Println(user)
	fmt.Println("----------------------------------------------------------")

	fmt.Println("--------------------------------------- first and assign ---------------------")

	user = User{}
	db.Where(User{Name: "shiksha"}).Assign(User{Age: "23"}).FirstOrInit(&user)
	fmt.Println(user)

	db.Where(User{Name: "Bimal"}).Assign(User{Age: "30"}).FirstOrInit(&user)
	fmt.Println(user)
	fmt.Println("----------------------------------------------------------")

	fmt.Println("--------------------------------------- first and assign ---------------------")

	user = User{}
	db.Where(User{Name: "shiksha"}).FirstOrCreate(&user)
	fmt.Println(user)

	user = User{}
	db.Where(User{Name: "Bimal"}).FirstOrCreate(&user)
	fmt.Println(user)
	fmt.Println("----------------------------------------------------------")

	fmt.Println("------------------------------- limit -----------------------------------------")
	users := []User{}
	db.Limit(3).Find(&users)
	fmt.Println(users)

	fmt.Println("--------------------------------------")

	fmt.Println("------------------------------- limit -----------------------------------------")
	users = []User{}
	count := 0
	db.Find(&users).Count(&count)
	fmt.Println(users)
	fmt.Println(count)
	fmt.Println("--------------------------------------")

	fmt.Println("------------------------------- Pluck -----------------------------------------")
	var ages []string
	db.Find(&users).Pluck("age", &ages)
	fmt.Println(ages)
	fmt.Println(users)
	fmt.Println("--------------------------------------")

	fmt.Println("end of code --------------------------------------")

}
