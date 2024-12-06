#pragma once
#include <WiFi.h>

#define WIFI_TIMEOUT_MS 20000

#define message_prefix "Wifi: "

class WiFiHandler
{
public:
	WiFiHandler()
	{
	}
	void connect(const std::string &ssid, const std::string &password)
	{
		m_ssid = ssid;
		m_password = password;

		Serial.print(String(message_prefix) + "Trying to connect with " + String(m_ssid.c_str()) + "\n");

		WiFi.begin(m_ssid.c_str(), m_password.c_str());

		unsigned long startAttemptTime = millis();

		while(WiFi.status() != WL_CONNECTED && millis() - startAttemptTime < WIFI_TIMEOUT_MS)
		{
		}

		if(WiFi.status() != WL_CONNECTED)
		{
			Serial.println(String(message_prefix) + "Failed to connect with " + String(m_ssid.c_str()) + "\n");
			return;
		}

		Serial.println(String(message_prefix) + "Connected with " + String(m_ssid.c_str()) + "\n");
		configTime(0, 0, "pool.ntp.org", "time.nist.gov");
	}
	bool isConnected()
	{
		return WiFi.status() == WL_CONNECTED;
	}
	std::string getSSID()
	{
		return WiFi.SSID().c_str();
	}

	std::string getMacAdress()
	{
		std::string mac = WiFi.macAddress().c_str();
		mac.erase(std::remove(mac.begin(), mac.end(), ':'), mac.end());
		return mac;
	}

private:
	std::string m_ssid;
	std::string m_password;
};