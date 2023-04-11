package scripts

import (
	"Nice-Things-Backend/initializers"
	"Nice-Things-Backend/models"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func init () {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func Seed () {
	log.Println("Seeding assistdirectorstamu@gmail.com into User table...")

	hash, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("PASSWORD")), 10)

	if err != nil {
		log.Fatal("Could not seed assistdirectorstamu@gmail.com into User table")
		return
	}

	user := models.User{
		Email: "assistdirectorstamu@gmail.com", 
		FirstName: "ASSIST", 
		LastName: "ADMIN", 
		Password: string(hash), 
		Admin: true,
	}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		log.Fatal("Could not seed assistdirectorstamu@gmail.com into User table")
		return
	}

	log.Println("Done seeding assistdirectorstamu@gmail.com into User table")
}