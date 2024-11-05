#include "MQTT.h"

#define message_prefix "MQQT: "

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
		Serial.print(String(message_prefix) + "Connecting with " + String(mqttServer) + "...\n");
		if(client.connect("ESP32Client", mqttUser, mqttPassword))
		{
			Serial.println(String(message_prefix) + "Connected.\n");
		}
		else
		{
			Serial.print(String(message_prefix) + "Cannot connect to the server: ");
			Serial.print(client.state());
			Serial.println(" - Retrying in 5 seconds\n");
			delay(5000);
		}
	}
}

void MQTTClientHandler::subscribeTopic(const char *topic)
{
	client.subscribe(topic);
}

void MQTTClientHandler::publishData(const char *topic, const char *payload)
{
	if(client.connected())
	{
		client.publish(topic, payload);
		Serial.print(String(message_prefix) + "Message send: ");
		Serial.print(payload);
		Serial.print(" to topic: ");
		Serial.println(String(topic) + "\n");
	}
	else
	{
		Serial.println(String(message_prefix) + "Cannot send message to the server. \n");
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
