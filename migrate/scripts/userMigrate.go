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

func UserMigrateUp () {
	log.Println("Migrating User table...")
	initializers.DB.AutoMigrate(&models.User{})
	log.Println("Done migrating User table")
}

func UserMigrateDown () {
	log.Println("Dropping User table...")
	initializers.DB.Migrator().DropTable(&models.User{})
	log.Println("Done dropping User table")
}