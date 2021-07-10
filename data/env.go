package data

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load(".env.dev")
	if err != nil {
		log.Fatal(err)
	}
}

func getEnvByKey(key string) string {
	return os.Getenv(key)
}
