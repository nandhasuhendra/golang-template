package config

import (
	"golang-template/helper"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	err := godotenv.Load()
  helper.PanicIfError(err)
}
