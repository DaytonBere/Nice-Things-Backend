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

func UserMigrate () {
	log.Println("Migrating Comment table...")
	initializers.DB.AutoMigrate(&models.User{})
	log.Println("Done migrating Comment table")
}