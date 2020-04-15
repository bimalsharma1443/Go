package DbConnect

import "github.com/jinzhu/gorm"

func dbConnect() {
	// connect to database
	db, err := gorm.Open("mysql", "root:@/godb?charset=utf8&parseTime=True&loc=Local")

	// checking a err
	if err != nil {
		panic("failed to connect database")
	}

	// db close
	defer db.Close()
}
