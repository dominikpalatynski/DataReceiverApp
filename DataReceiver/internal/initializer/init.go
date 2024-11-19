package initializer

import (
	"data_receiver/internal/broker"
	"data_receiver/internal/config"
	"data_receiver/internal/device"
	"log"
)

func InitializeApplication() {

	config, err := config.LoadConfig()
	if err != nil {
        log.Fatalf("Fatal error during config load %v", err)
    }

	log.Print("Config loaded")

    deviceManager, err := device.NewDeviceManager(config.Database.Url, config.Database.Token, config.Database.Org)
	if err != nil {
        log.Fatalf("Fatal error during creating DeviceMenager %v", err)
    }

	log.Print("DeviceMenager created")

    mqttClient, err := broker.NewClient(config.Broker.BrokerUrl)
    if err != nil {
        log.Fatalf("Błąd konfiguracji MQTT: %v", err)
    }

    log.Println("Connected MQTT client")

    if token := mqttClient.Subscribe(config.Broker.TopicPattern, 0, deviceManager.ProcessMQTTMessage); token.Wait() && token.Error() != nil {
        log.Fatalf("Błąd subskrybowania: %v", token.Error())
    }

	log.Println("Subscribed on topic: %s", config.Broker.TopicPattern)
}