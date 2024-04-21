package config

import (
	"github.com/joho/godotenv"
)

func ConfigInit() (err error) {
	err = godotenv.Load(".env")

	return
}
