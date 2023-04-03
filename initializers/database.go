package initializers

import (
	"log"
	"os"
	"strings"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB 
var err error


func ConnectToDB () {	
	path := strings.Split(os.Args[0], "/")
	filename := path[len(path) - 1]
	log.Printf("Connecting to the DB in file: `%s.go`...\n", filename)
	dsn := os.Getenv("DB_URL")
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to DB")
	} else {
		log.Println("Connected to the DB")
	}
}