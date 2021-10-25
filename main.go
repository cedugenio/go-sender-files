package main

import (
	"cedugenio/go-sender-files/app"
	"cedugenio/go-sender-files/constants"
	"flag"
	"fmt"
	"os"
)

var (
	configPath string
)

func init() {
	flag.StringVar(&configPath, "properties", constants.CONFIG_FILE,
		" Informe o config path")
	flag.Parse()
}

func main() {
	err := app.Start(configPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
