package initializers

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func LoadEnviromentalVariables() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
		log.Fatal("Error loading .env file")

	}
}
