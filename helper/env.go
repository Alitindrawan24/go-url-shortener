package helper

import (
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

func GetEnvVariableKey(key string, alternative string) string {
	if os.Getenv("APP_ENV") == "testing" {
		err := godotenv.Load("../.env.testing")
		if err != nil {
			panic("Error loading .env.testing file")
		}
	} else {
		path_dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			panic("Error get path dir")
		}

		err = godotenv.Load(filepath.Join(path_dir, ".env"))
		if err != nil {
			panic("Error loading .env file")
		}
	}

	value := os.Getenv(key)

	if value == "" {
		return alternative
	}

	return value
}
