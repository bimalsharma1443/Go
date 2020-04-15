package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// user model
type User struct {
	gorm.Model
	Name  string `gorm:"type:varchar(100)"`
	Email string `gorm:"type:varchar(100)"`
	Phone string `gorm:"type:varchar(10)"`
}

func main() {
	fmt.Println("MUX with ORM")

	// connecting to a db
	db := dbConnect()

	defer db.Close()

	// migrating a data
	migration(db)

	// handling arequest
	handelRequest()
}

func migration(db *gorm.DB) {
	fmt.Println("migrationg a data")
	db.AutoMigrate(&User{})
}

func dbConnect() *gorm.DB {
	db, err := gorm.Open("mysql", "root:@/godb?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("unable to connect a db : ", err)
	}

	return db
}

func handelRequest() {
	fmt.Println("handling a request")

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.HandleFunc("/", allUser).Methods("GET")
	myRouter.HandleFunc("/user/{id}", viewUser).Methods("GET")
	myRouter.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	myRouter.HandleFunc("/user/add", addUser).Methods("POST")
	myRouter.HandleFunc("/user/edit/{id}", editUser).Methods("POST")

	log.Fatal(http.ListenAndServe(":8081", myRouter))
}

func allUser(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()

	users := []User{}

	db.Find(&users)

	for _, v := range users {
		fmt.Fprintln(w, v)
	}
}

func viewUser(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()

	user := User{}

	db.Find(&user)

	fmt.Fprintln(w, user)

}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	db := dbConnect()

	user := User{}

	vars := mux.Vars(r)

	db.Find(&user, vars["id"])

	db.Delete(&user)

	fmt.Fprintf(w, "user is deleted")
}

func addUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := dbConnect()

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	json.NewEncoder(w).Encode(&user)

	db.Create(&user)

}

func editUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	db := dbConnect()

	vars := mux.Vars(r)

	var user User

	db.Find(&user, vars["id"])

	_ = json.NewDecoder(r.Body).Decode(&user)
	json.NewEncoder(w).Encode(&user)

	db.Save(&user)

}
