#include "MQTT.h"

MQTTClientHandler::MQTTClientHandler(const char *mqttServer, int mqttPort, const char *mqttUser,
									 const char *mqttPassword)
	: mqttServer(mqttServer), mqttPort(mqttPort), mqttUser(mqttUser), mqttPassword(mqttPassword), client(espClient)
{
	client.setServer(mqttServer, mqttPort);
}

void MQTTClientHandler::begin()
{
	connectToMQTT();
}

void MQTTClientHandler::setCallback(MQTT_CALLBACK_SIGNATURE)
{
	client.setCallback(callback);
}

void MQTTClientHandler::connectToMQTT()
{
	while(!client.connected())
	{
		Serial.print("Łączenie z brokerem MQTT...");
		if(client.connect("ESP32Client", mqttUser, mqttPassword))
		{
			Serial.println("połączono.");
			client.subscribe("test/topic");
		}
		else
		{
			Serial.print("Błąd połączenia: ");
			Serial.print(client.state());
			Serial.println(" - Próba ponownie za 5 sekund");
			delay(5000);
		}
	}
}

void MQTTClientHandler::publishData(const char *topic, const char *payload)
{
	if(client.connected())
	{
		client.publish(topic, payload);
		Serial.print("Wysłano wiadomość: ");
		Serial.print(payload);
		Serial.print(" na temat: ");
		Serial.println(topic);
	}
	else
	{
		Serial.println("Nie można wysłać - brak połączenia z MQTT");
	}
}

bool MQTTClientHandler::isConnected()
{
	return client.connected();
}

void MQTTClientHandler::loop()
{
	client.loop();
}
