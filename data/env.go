package data

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load("../app.env.development.local")
	if err != nil {
		log.Fatal(err)
	}
}

func getEnvByKey(key string) string {
	return os.Getenv(key)
}
