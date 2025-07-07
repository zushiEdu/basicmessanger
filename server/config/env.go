package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Env struct {
	DBHost  string
	DBName  string
	DBPort  string
	DBUser  string
	DBPass  string
	MODE    string
	APIHost string
	APIPort string
}

var env *Env

func LoadEnv() *Env {
	fmt.Println("Loading environment variables")
	godotenv.Load(".env")
	env = &Env{DBHost: os.Getenv("DB_SOURCE"),
		DBName:  os.Getenv("DB_NAME"),
		DBPort:  os.Getenv("DB_PORT"),
		DBUser:  os.Getenv("DB_USER"),
		DBPass:  os.Getenv("DB_PASS"),
		MODE:    os.Getenv("MODE"),
		APIHost: os.Getenv("API_HOST"),
		APIPort: os.Getenv("API_PORT"),
	}
	return env
}

func GetEnv() *Env {
	return env
}
