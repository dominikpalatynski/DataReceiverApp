#ifndef MQTTCLIENTHANDLER_H
#define MQTTCLIENTHANDLER_H

#include <PubSubClient.h>
#include <WiFiClient.h>

class MQTTClientHandler
{
public:
	MQTTClientHandler(const char *mqttServer, int mqttPort, const char *mqttUser, const char *mqttPassword);

	void begin();
	void setCallback(MQTT_CALLBACK_SIGNATURE);
	void connectToMQTT();
	void subscribeTopic(const char *topic);
	void publishData(const char *topic, const char *payload);
	bool isConnected();
	void loop();

private:
	const char *mqttServer;
	int mqttPort;
	const char *mqttUser;
	const char *mqttPassword;

	WiFiClient espClient;
	PubSubClient client;
};

#endif
