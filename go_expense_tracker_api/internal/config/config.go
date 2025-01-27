package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Load() map[string]string {
	var env string = os.Getenv("APP_ENV")

	if env == "" {
		env = "local"
	}

	myEnv, err := godotenv.Read(".env." + env)
	if err != nil {
		fmt.Println(err)
		panic("failed to load env")
	}

	return myEnv
}
