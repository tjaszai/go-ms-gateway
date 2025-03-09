package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

func Config(key string, defaultValue ...string) string {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	value, ok := os.LookupEnv(key)
	if (!ok || value == "") && len(defaultValue) > 0 {
		fmt.Printf("Using default value '%s'\n", defaultValue[0])
		return defaultValue[0]
	}

	return value
}
