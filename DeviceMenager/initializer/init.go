package initializer

import (
	"ConfigApp/cache"
	"ConfigApp/config"
	"ConfigApp/logging"
	"ConfigApp/server"
	"ConfigApp/storage"
	"ConfigApp/user"

	supa "github.com/nedpals/supabase-go"
)

func InitializeApplication() *server.APIServer {
	config, err := config.LoadConfig()
	if err != nil {
		logging.Log.Fatalf("Cannot load configuration: %v", err)
		return nil
	}

	redisClient, err := cache.NewRedisClient(*config)
	
	if err != nil {
		logging.Log.Fatalf("Cannot establish connection with Redis", err)
	}

	client := supa.CreateClient(config.Database.Url, config.Database.Key)
	storage := storage.NewSupabaseStorage(client)
	userHandler := user.NewSupabaseUserHandler(client)
	server := server.NewAPIServer(storage, userHandler, *config, redisClient)

	return server
}