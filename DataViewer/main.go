package main

import (
	"data_viewer/config"
	"data_viewer/server"
	"data_viewer/storage"
	"log"
)

func main() {

	
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	client := storage.NewClient(config.Database.Url, config.Database.Key, config.Database.Org)

	server := server.NewAPIServer(client, *config)
	server.Run()
	select{}
}