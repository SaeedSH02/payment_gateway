package config


import (
	"log"

	"github.com/joho/godotenv"
)



func init() {

	// ctx, cancel := context.WithTimeout(context.Background(), 1*time.Minute)
	// defer cancel()

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Can't Read .env")
	}


}
