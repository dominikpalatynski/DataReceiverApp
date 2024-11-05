#include "WiFiHandler.h"

#define WIFI_TIMEOUT_MS 20000

#define message_prefix "Wifi: "

WiFiHandler::WiFiHandler(const char *ssid, const char *password) : m_ssid(ssid), m_password(password)
{
}

void WiFiHandler::connect()
{
	Serial.print(String(message_prefix) + "Trying to connect with " + String(m_ssid) + "\n");

	WiFi.begin(m_ssid, m_password);

	unsigned long startAttemptTime = millis();

	while(WiFi.status() != WL_CONNECTED && millis() - startAttemptTime < WIFI_TIMEOUT_MS)
	{
	}

	if(WiFi.status() != WL_CONNECTED)
	{
		Serial.println(String(message_prefix) + "Failed to connect with " + String(m_ssid) + "\n");
		return;
	}

	Serial.println(String(message_prefix) + "Connected with " + String(m_ssid) + "\n");
}

bool WiFiHandler::isConnected()
{
	return WiFi.status() == WL_CONNECTED;
}

String WiFiHandler::getSSID()
{
	return WiFi.SSID();
}