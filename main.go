package main

import (
	"fmt"

	"./config"
)

func main() {
	config, _ := config.PrepareConfig()
	fmt.Println(config)
}
