#include "WiFiHandler.h"

WiFiHandler::WiFiHandler(const char *ssid, const char *password) : m_ssid(ssid), m_password(password)
{
}

void WiFiHandler::connect()
{
	Serial.print("WiFi: Trying to connect with " + String(m_ssid) + "\n");

	WiFi.begin(m_ssid, m_password);

	Serial.println("Wifi: Connected with" + String(m_ssid) + "\n");
}

bool WiFiHandler::isConnected()
{
	return WiFi.status() == WL_CONNECTED;
}

String WiFiHandler::getSSID()
{
	return WiFi.SSID();
}