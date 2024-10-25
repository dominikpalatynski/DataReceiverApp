#include "IBusHandler.h"
#include "MQTT.h"
#include "WiFiHandler.h"
#include "display.h"

const char *ssid = "Niusia_Internet";
const char *password = "Kapibara1337!";
const char *mqttServer = "74.248.137.150";
const int mqttPort = 1883;
const char *mqttUser = "mqttUser";
const char *mqttPassword = "mqttPass";
const char *topic = "devices/device1/measurements";
const char *topic2 = "devices/device2/measurements";

WiFiHandler wifiHandler(ssid, password);
MQTTClientHandler mqttHandler(mqttServer, mqttPort, mqttUser, mqttPassword);
OLEDDisplay myDisplay;
HardwareSerial mySerial(2);
IBusHandler ibusHandler(mySerial);

void mqttCallback(char *topic, byte *message, unsigned int length)
{
	Serial.print("Otrzymano wiadomość na temacie: ");
	Serial.println(topic);
	Serial.print("Wiadomość: ");
	for(int i = 0; i < length; i++)
	{
		Serial.print((char) message[i]);
	}
	Serial.println();
}

int liczba = 0;

void setup()
{
	Serial.begin(115200);
	wifiHandler.connect();
	myDisplay.begin();
	mqttHandler.begin();
	mqttHandler.setCallback(mqttCallback);
	ibusHandler.begin();
}

void loop()
{
	delay(2000);

	liczba += 69;

	if(liczba > 10000)
	{
		liczba = 0;
	}
	ibusHandler.update();
	mqttHandler.loop();

	if(mqttHandler.isConnected())
	{
		mqttHandler.publishData(topic, String(liczba).c_str());
		mqttHandler.publishData(topic, String(10000 - liczba).c_str());
	}

	myDisplay.setRow(0, "Wifi:");
	myDisplay.setRow(1, wifiHandler.getSSID());
	myDisplay.setRow(2, "MQQT topic:");
	myDisplay.setRow(3, String(topic).substring(0, 21));
	myDisplay.setRow(4, "MQQT data: ");
	myDisplay.setRow(5, String(liczba));
}
