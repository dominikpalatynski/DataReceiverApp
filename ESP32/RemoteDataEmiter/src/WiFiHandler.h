#ifndef WIFIHANDLER_H
#define WIFIHANDLER_H

#include <WiFi.h>

class WiFiHandler
{
public:
	WiFiHandler(const char *ssid, const char *password);
	void connect();
	bool isConnected();
	String getSSID();

private:
	const char *m_ssid;
	const char *m_password;
};

#endif
