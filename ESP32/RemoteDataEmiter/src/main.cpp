#include "IBusHandler.h"
#include "MQTT.h"
#include "WiFiHandler.h"
#include "display.h"

#include <freertos/FreeRTOS.h>
#include <freertos/queue.h>
#include <freertos/task.h>

QueueHandle_t wifiStatusQueue;

void displayTask(void *pvParameters);
void wifiTask(void *pvParameters);
void mqttTask(void *pvParameters);

const char *ssid = "Niusia_Internet";
const char *password = "Kapibara1337!";
const char *mqttServer = "74.248.137.150";
const int mqttPort = 1883;
const char *mqttUser = "mqttUser";
const char *mqttPassword = "mqttPass";
const char *topic = "devices/device1/measurements";

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

	wifiStatusQueue = xQueueCreate(1, sizeof(bool));

	if(wifiStatusQueue == NULL)
	{
		Serial.println("Błąd przy tworzeniu kolejki.");
		return;
	}

	xTaskCreate(displayTask, "Display Task", 4096, NULL, 1, NULL);
	xTaskCreate(wifiTask, "WiFi Task", 4096, NULL, 1, NULL);
	xTaskCreate(mqttTask, "MQTT Task", 4096, NULL, 1, NULL);
}

void loop()
{
}

void displayTask(void *pvParameters)
{
	bool wifiConnected = false;

	myDisplay.begin();
	myDisplay.setRow(0, "Wifi:");
	myDisplay.setRow(1, "not connected");
	myDisplay.setRow(2, "MQTT topic:");
	myDisplay.setRow(4, "MQTT data:");

	while(true)
	{
		xQueuePeek(wifiStatusQueue, &wifiConnected, portMAX_DELAY);

		if(wifiConnected)
		{
			myDisplay.setRow(1, wifiHandler.getSSID());
		}
		else
		{
			myDisplay.setRow(1, "not connected");
		}

		vTaskDelay(1000 / portTICK_PERIOD_MS);
	}
}

void wifiTask(void *pvParameters)
{
	bool wifiConnected = false;
	while(true)
	{
		if(wifiHandler.isConnected())
		{
			wifiConnected = true;
			xQueueOverwrite(wifiStatusQueue, &wifiConnected);
			vTaskDelay(5000 / portTICK_PERIOD_MS);
			continue;
		}
		wifiConnected = false;
		xQueueOverwrite(wifiStatusQueue, &wifiConnected);
		wifiHandler.connect();
	}
}

void publishData()
{
	liczba += 69;

	if(liczba > 10000)
	{
		liczba = 0;
	}
	mqttHandler.loop();

	if(mqttHandler.isConnected())
	{
		mqttHandler.publishData(topic, String(liczba).c_str());
		myDisplay.setRow(3, String(topic).substring(0, 20));
		myDisplay.setRow(5, String(liczba));
	}
}

void mqttTask(void *pvParameters)
{
	bool wifiConnected = false;
	bool wasConnected = false;

	while(true)
	{
		xQueuePeek(wifiStatusQueue, &wifiConnected, portMAX_DELAY);
		if(wifiConnected)
		{
			if(!wasConnected)
			{
				mqttHandler.begin();
				mqttHandler.setCallback(mqttCallback);

				wasConnected = true;
			}

			publishData();
		}
		else
		{
			wasConnected = false;
		}

		vTaskDelay(5000 / portTICK_PERIOD_MS);
	}
}
