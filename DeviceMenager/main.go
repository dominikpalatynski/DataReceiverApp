package main

import (
	"ConfigApp/server"
	"ConfigApp/storage"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	if err := godotenv.Load(".env"); err !=nil {
		log.Fatal("Error loading .env")
	}
}

func main() {
	LoadEnv()

	storage := storage.NewSupabaseStorage(os.Getenv("SUPABASE_URL"), os.Getenv("SUPABASE_KEY"))
	server := server.NewAPIServer(storage)
	server.Run()
	select{}
}