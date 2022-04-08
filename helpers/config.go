package helpers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func Load(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalln(err)
	}

	return os.Getenv(key)
}
