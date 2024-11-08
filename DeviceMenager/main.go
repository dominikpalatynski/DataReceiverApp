package main

import (
	"ConfigApp/server"
	"ConfigApp/storage"
	"ConfigApp/user"
	"log"
	"os"

	"github.com/joho/godotenv"
	supa "github.com/nedpals/supabase-go"
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err !=nil {
		log.Fatal("Error loading .env")
	}
}

func main() {
	LoadEnv()

	client := supa.CreateClient(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))

	storage := storage.NewSupabaseStorage(client)
	userHandler := user.NewSupabaseUserHandler(client)
	server := server.NewAPIServer(storage, userHandler)
	server.Run()
	select{}
}