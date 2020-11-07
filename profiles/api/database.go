package api

import (
	"database/sql"
	"log"
	"time"

	//MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() *sql.DB {
	log.Println("attempting connections")
	var err error
	DB, err = sql.Open("mysql", "root:root@tcp(172.28.1.2:3306)/profiles")
	if err != nil {
		log.Println("error opening DB connection")
		panic(err.Error())
	}

	err = DB.Ping()
	for err != nil {
		log.Println("couldnt connect, waiting 20 seconds before retrying")
		time.Sleep(20*time.Second)
		// Connect again, use the same connection function as you did above ^
		err = DB.Ping()
	}

	return DB
}
