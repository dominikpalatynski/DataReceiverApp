package initializer

import (
	"ConfigApp/cache"
	"ConfigApp/config"
	"ConfigApp/server"
	"ConfigApp/storage"
	"ConfigApp/user"
	"log"

	supa "github.com/nedpals/supabase-go"
)

func InitializeApplication() *server.APIServer {
	config, err := config.LoadConfig()
	if err != nil {
		log.Panic("Panic error during config load", err)
	}

	redisClient, err := cache.NewRedisClient(*config)
	
	if err != nil {
		log.Panic("Cannot establish connection with Redis", err)
	}

	client := supa.CreateClient(config.Database.Url, config.Database.Key)
	storage := storage.NewSupabaseStorage(client)
	userHandler := user.NewSupabaseUserHandler(client)
	server := server.NewAPIServer(storage, userHandler, *config, redisClient)

	return server
}