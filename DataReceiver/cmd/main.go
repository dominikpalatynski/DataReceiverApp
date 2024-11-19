package main

import "data_receiver/internal/initializer"

func main() {
	initializer.InitializeApplication()
	select {}
}