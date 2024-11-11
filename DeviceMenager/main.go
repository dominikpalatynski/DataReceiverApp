package main

import (
	"ConfigApp/config"
	"ConfigApp/server"
	"ConfigApp/storage"
	"ConfigApp/user"
	"log"

	supa "github.com/nedpals/supabase-go"
)

func main() {

	
	config, err := config.LoadConfig()

	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	client := supa.CreateClient(config.Database.Url, config.Database.Key)

	storage := storage.NewSupabaseStorage(client)
	userHandler := user.NewSupabaseUserHandler(client)
	server := server.NewAPIServer(storage, userHandler, *config)
	server.Run()
	select{}
}