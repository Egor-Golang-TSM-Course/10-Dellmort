package main

import (
	"lesson10/config"
	"lesson10/processor"
	"log"
)

func main() {
	config, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	buffer := make([]byte, 1024)
	analysis, err := processor.New(config, buffer)
	if err != nil {
		log.Fatal(err)
	}

	err = analysis.Analysis()
	if err != nil {
		log.Fatal(err)
	}
}
