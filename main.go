package main

import (
	"fmt"
	"go-trial-class/cli"
	"go-trial-class/config"
)

func main() {
	fmt.Println("Hello World")
	config.DBConnect()
	cli.MainMenu()
}
