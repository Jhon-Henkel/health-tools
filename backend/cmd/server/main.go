package main

import (
	"github.com/Jhon-Henkel/health-tools/tree/main/backend/configs"
)

func main() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		panic(err)
	}
	println(config.DbDriver)
}
