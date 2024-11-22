#include "IBusHandler.h"
#include "JsonGenerator.h"
#include "MQTT.h"
#include "WiFiHandler.h"
#include "display.h"

#include <FS.h>
#include <HTTPClient.h>
#include <SPIFFS.h>
#include <freertos/FreeRTOS.h>
#include <freertos/queue.h>
#include <freertos/task.h>

QueueHandle_t wifiStatusQueue;
SemaphoreHandle_t xSemaphore;

void displayTask(void *pvParameters);
void wifiTask(void *pvParameters);
void mqttTask(void *pvParameters);

const char *ssid = "Niusia_Internet";
const char *password = "Kapibara1337!";
const char *mqttServer = "74.248.137.150";
const int mqttPort = 1883;
const char *mqttUser = "mqttUser";
const char *mqttPassword = "mqttPass";
String topic;
const String serverUrl = "http://example.com/api/device-id"; // URL API
const char *filePath = "/device_id.txt";

WiFiHandler wifiHandler(ssid, password);
MQTTClientHandler mqttHandler(mqttServer, mqttPort, mqttUser, mqttPassword);
OLEDDisplay myDisplay;
HardwareSerial mySerial(2);
IBusHandler ibusHandler(mySerial);

void setupSPIFFS()
{
	if(!SPIFFS.begin(true))
	{
		Serial.println("Failed to initialize SPIFFS");
		return;
	}
}

String getMacAddress()
{
	String mac = WiFi.macAddress();
	mac.replace(":", "");
	return mac;
}

String sendHttpRequestWithMAC(const String &mac)
{
	HTTPClient http;
	http.begin(serverUrl);
	http.addHeader("Content-Type", "application/json");

	JsonServerRequestID jsonGen;

	String token = "sample_token";
	jsonGen.initializeJson(mac, token);
	String json = jsonGen.generateJsonString();

	Serial.println("Request payload: " + json);

	int httpCode = http.POST(json);
	if(httpCode > 0)
	{
		if(httpCode == HTTP_CODE_OK)
		{
			String response = http.getString();
			http.end();
			return response;
		}
	}
	else
	{
		Serial.println("HTTP request failed");
	}

	http.end();
	return "";
}

void saveToFile(const char *path, const String &data)
{
	File file = SPIFFS.open(path, FILE_WRITE);
	if(!file)
	{
		Serial.println("Failed to open file for writing");
		return;
	}
	file.print(data);
	file.close();
	Serial.println("Data saved to file.");
}

String readFromFile(const char *path)
{
	File file = SPIFFS.open(path, FILE_READ);
	if(!file)
	{
		Serial.println("Failed to open file for reading");
		return "";
	}

	String content = file.readString();
	file.close();
	return content;
}

void mqttCallback(char *topic, byte *message, unsigned int length)
{
	Serial.print("Recieved message on topic: ");
	Serial.println(topic);
	Serial.print("message: ");
	for(int i = 0; i < length; i++)
	{
		Serial.print((char) message[i]);
	}
	Serial.println();
}

void publishData()
{
	int liczba = 0;
	liczba += 69;

	if(liczba > 10000)
	{
		liczba = 0;
	}
	mqttHandler.loop();

	if(mqttHandler.isConnected())
	{
		JsonSensorDataGenerator jsonGen;
		jsonGen.initializeJson("device1");
		jsonGen.addSensorData(liczba, "device1sensor1");

		String jsonString = jsonGen.generateJsonString();
		mqttHandler.publishData(topic.c_str(), jsonString.c_str());
		myDisplay.setRow(3, String(topic).substring(0, 20));
		myDisplay.setRow(5, String(liczba));
	}
}

void setup()
{
	Serial.begin(115200);

	setupSPIFFS();

	xSemaphore = xSemaphoreCreateBinary();

	if(xSemaphore == NULL)
	{
		Serial.println("Can't create semaphore");
		return;
	}
	wifiStatusQueue = xQueueCreate(1, sizeof(bool));

	if(wifiStatusQueue == NULL)
	{
		Serial.println("Can't create wifiStatus queue");
		return;
	}

	xTaskCreate(displayTask, "Display Task", 4096, NULL, 1, NULL);
	xTaskCreate(wifiTask, "WiFi Task", 4096, NULL, 1, NULL);
	xTaskCreate(getDeviceIdTask, "Device ID request Task", 4096, NULL, 1, NULL);
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

void getDeviceIdTask(void *pvParameters)
{
	String id;
	while(id.isEmpty())
	{
		if(SPIFFS.exists(filePath))
		{
			id = readFromFile(filePath);
			if(id.length() > 0)
			{
				Serial.println("ID found in file: " + id);
				xSemaphoreGive(xSemaphore);
				topic = "devices/" + id + "/measurments";
				return;
			}
		}

		Serial.println("ID not found. Sending HTTP request...");
		String mac = getMacAddress();
		id = sendHttpRequestWithMAC(mac);
		if(id.length() > 0)
		{
			saveToFile(filePath, id);
			xSemaphoreGive(xSemaphore);
			topic = "devices/" + id + "/measurments";
			return;
		}
		vTaskDelay(5000 / portTICK_PERIOD_MS);
	}

	Serial.println("Failed to get ID from server");
}

void mqttTask(void *pvParameters)
{
	bool wifiConnected = false;
	bool wasConnected = false;

	while(true)
	{
		if(xSemaphoreTake(xSemaphore, portMAX_DELAY) == pdTRUE)
		{
			xQueuePeek(wifiStatusQueue, &wifiConnected, portMAX_DELAY);
			if(wifiConnected)
			{
				if(!mqttHandler.isConnected())
				{
					mqttHandler.begin();
					mqttHandler.setCallback(mqttCallback);
				}

				publishData();
			}

			vTaskDelay(5000 / portTICK_PERIOD_MS);
		}
	}
}
