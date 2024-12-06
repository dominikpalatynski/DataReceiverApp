#ifndef MQTTCLIENTHANDLER_H
#define MQTTCLIENTHANDLER_H

#include <PubSubClient.h>
#include <WiFiClient.h>
#include <string>

#define MESSAGE_PREFIX "MQTT: "

class MQTTClientHandler
{
public:
	MQTTClientHandler()
	{
	}

	void begin(const std::string &server, int port, const std::string &user, const std::string &password)
	{
		mqttServer = server;
		mqttPort = port;
		mqttUser = user;
		mqttPassword = password;
		client.setServer(mqttServer.c_str(), mqttPort);
		connectToMQTT();
	}

	void setCallback(MQTT_CALLBACK_SIGNATURE)
	{
		client.setCallback(callback);
	}

	void connectToMQTT()
	{
		while(!client.connected())
		{
			Serial.print(String(MESSAGE_PREFIX) + "Connecting to " + mqttServer.c_str() + "...\n");
			if(client.connect("ESP32Client", mqttUser.c_str(), mqttPassword.c_str()))
			{
				Serial.println(String(MESSAGE_PREFIX) + "Connected.\n");
			}
			else
			{
				Serial.print(String(MESSAGE_PREFIX) + "Cannot connect to the server: ");
				Serial.print(client.state());
				Serial.println(" - Retrying in 5 seconds\n");
				delay(5000);
			}
		}
	}

	void subscribeTopic(const std::string &topic)
	{
		client.subscribe(topic.c_str());
	}

	void publishData(const std::string &topic, const std::string &payload)
	{
		if(client.connected())
		{
			client.publish(topic.c_str(), payload.c_str());
			Serial.print(String(MESSAGE_PREFIX) + "Message sent: ");
			Serial.print(payload.c_str());
			Serial.print(" to topic: ");
			Serial.println(topic.c_str());
		}
		else
		{
			Serial.println(String(MESSAGE_PREFIX) + "Cannot send message to the server.\n");
		}
	}

	bool isConnected()
	{
		return client.connected();
	}

	void loop()
	{
		client.loop();
	}

private:
	std::string mqttServer;
	int mqttPort;
	std::string mqttUser;
	std::string mqttPassword;

	WiFiClient espClient;
	PubSubClient client;
};

#endif
