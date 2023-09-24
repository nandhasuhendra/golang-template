package main

import (
	"fmt"
	"golang-template/config"
)

func main() {
	fmt.Println("My Personal Template For Golang Project")
	config.LoadEnv()
  config.NewDB()
}
