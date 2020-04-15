package main

import (
	"encoding/csv"
	"fmt"
	// "io"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

type userLog struct {
	user     string
	is_login string
	reason   string
	in_time  string
	ip       string
}

func main() {

	fileChan := make(chan userLog)

	// go routine readfile
	go readfile(fileChan)

	// insert into db
	inserIntoDb(fileChan)
}

func readfile(fileChan chan<- userLog) {

	// open csv file
	file, err := os.Open("user_logs.csv")

	if err != nil {
		log.Fatalln("Could not open the file", err)
	}

	r := csv.NewReader(file)

	for {
		// reading a file
		record, err := r.Read()

		if err != nil {
			log.Fatalln(err)
			break
		}

		// create user struct
		user := userLog{
			user:     record[1],
			is_login: record[2],
			reason:   record[3],
			in_time:  record[4],
			ip:       record[5],
		}

		fileChan <- user
	}

	close(fileChan)
}

func inserIntoDb(fileChan <-chan userLog) {

	// create db connection
	fmt.Println("connection to db...")
	db, err := sql.Open("mysql", "root:@/godemo")

	if err != nil {
		panic(err.Error())
		return
	}

	// close db connection
	defer db.Close()

	// creating a statement
	fmt.Println("creating statement")
	stmt, err := db.Prepare("INSERT INTO user_logs(user_id, is_login, reason, in_time, ip) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Fatalln(err)
	}

	// close statement
	defer stmt.Close()

	// inserrting a data into db
	for v := range fileChan {
		_, err := stmt.Exec(v.user, v.is_login, v.reason, v.in_time, v.ip)

		if err != nil {
			log.Fatalln(err)
		}
	}

	fmt.Println("exit")
}
