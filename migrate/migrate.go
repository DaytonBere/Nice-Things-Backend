package main

import (
	"Nice-Things-Backend/migrate/scripts"
	"log"
)

func main () {
	log.Println("Starting migration scripts...")
	scripts.UserMigrate()
	scripts.NiceThingMigrate()
	log.Println("Done with migration scripts")
}