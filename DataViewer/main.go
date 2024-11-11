package main

import (
	"data_viewer/config"
	"data_viewer/server"
	"log"
)

func main() {

	
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}


	server := server.NewAPIServer(*config)
	server.Run()
	select{}
}