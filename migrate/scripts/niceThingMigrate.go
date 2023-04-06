package scripts

import (
	"Nice-Things-Backend/initializers"
	"Nice-Things-Backend/models"
	"log"
)

func init () {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func NiceThingMigrateUp () {
	log.Println("Migrating NiceThing table...")
	initializers.DB.AutoMigrate(&models.NiceThing{})
	log.Println("Done migrating NiceThing table")
}

func NiceThingMigrateDown () {
	log.Println("Dropping NiceThing table...")
	initializers.DB.Migrator().DropTable(&models.NiceThing{})
	log.Println("Done dropping NiceThing table")
}
