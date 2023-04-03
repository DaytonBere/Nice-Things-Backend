package initializers

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnvVariables () {
	path := strings.Split(os.Args[0], "/")
	filename := path[len(path) - 1]
	log.Printf("Loading .env file in file: `%s.go`...\n", filename)
	err := godotenv.Load();

    if err != nil {
        log.Fatal("Error loading .env file")
    } else {
		log.Println("Loaded .env file")
	}
}