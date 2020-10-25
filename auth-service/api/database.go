package api

import (
	"database/sql"
	"log"
	"time"

	//MySQL driver
	_ "github.com/go-sql-driver/mysql"
)

//DB represents the connection to the MySQL database
var (
	DB *sql.DB
)

//InitDB creates the MySQL database connection
func InitDB() *sql.DB {

	log.Println("attempting connections")

	var err error

	// Open a SQL connection to the docker container hosting the database server
	// Assign the connection to the "DB" variable
	// Look at how it's done in the other microservices!
	dbType := "mysql"
	username := "root"
	password := "root"
	ipAddress := "tcp(172.28.1.2:3306)"
	dbName := "/auth"
	// "YOUR CODE HERE"
	DB, err = sql.Open(dbType, username+":"+password+"@"+ipAddress+dbName)
	if err != nil {
		log.Println("error opening DB connection")
		panic(err.Error())
	}

	err = DB.Ping()
	for err != nil {
		log.Println("couldnt connect, waiting 10 seconds before retrying")
		time.Sleep(10*time.Second)
		// Connect again, use the same connection function as you did above ^
		err = DB.Ping()
	}

	return DB
}
