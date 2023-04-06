package main

import (
	"Nice-Things-Backend/migrate/scripts"
	"fmt"
	"log"
)

func main () {
	fmt.Println("Enter up or down for migration: ")

	var direction string
	fmt.Scanln(&direction)


	if direction == "up" {
		log.Println("Starting migration-up scripts...")
		scripts.UserMigrateUp()
		scripts.NiceThingMigrateUp()
		scripts.Seed()
		log.Println("Done with migration-up scripts")
	} else if direction == "down" {
		log.Println("Starting migration-down scripts...")
		scripts.UserMigrateDown()
		scripts.NiceThingMigrateDown()
		log.Println("Done with migration-down scripts")
	} else {
		log.Fatal("Value is neither 'up' or 'down'")
	}
}