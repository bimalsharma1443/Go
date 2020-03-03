package main

import "fmt"
import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "encoding/json"


type User struct {
	id     	int
	name    string
	age   	int
	mobile 	int
}


func main() {

	var user User
	var users []User

	db, err := sql.Open("mysql", "root:@/godemo")

	if err != nil {
    	panic(err.Error()) // Just for example purpose. You should use proper error handling instead of panic
		return
	}
	
	defer db.Close()

	err = db.Ping()
	if err != nil {
	    panic(err.Error()) // proper error handling instead of panic in your app
	    return
	}


	rows,err := db.Query("select * from user")

	if err != nil {
		fmt.Println(err);
		return
	}

	for rows.Next() {
		rows.Scan(&user.id,&user.name,&user.age,&user.mobile)
		users = append(users, user)
	}

	defer rows.Close()

	jsonResponse, jsonError := json.Marshal(users)
	if jsonError != nil {
		fmt.Println(jsonError)
		return
		//returnErrorResponse(response, request)
	}

	fmt.Println(string(jsonResponse));

}