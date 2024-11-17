package main

import (
	"ConfigApp/initializer"
)

func main() {
	server := initializer.InitializeApplication()
	server.Run()
	select{}
}