package main

import (
	"DeviceStateHandler/internal/config"
	"DeviceStateHandler/internal/mqtt"
	"DeviceStateHandler/internal/storage"
	"log"
)

func main() {

    config, err := config.LoadConfig()

    if err != nil {
        log.Fatal("cannot load config")
    }

    redisClient, err := storage.NewRedisClient(config.Cache.Url, config.Cache.Password)

    if err != nil {
        log.Fatal("cannot connect to redis")
    }

    influxDbClient := storage.NewClient(config.Database.Url, config.Database.Token, config.Database.Org)


    mqttClient, err := mqtt.NewMqttClient(config.Broker.BrokerUrl, redisClient, influxDbClient, config.Server.DeviceMenagerUrl)
    if err != nil {
        log.Fatalf("Błąd konfiguracji MQTT: %v", err)
    }

    log.Println("Connected MQTT client")

    if err := mqttClient.Subcribe(config.Broker.StatusTopicPattern); err != nil {
        log.Fatal(err.Error())
    }
	select {}
}