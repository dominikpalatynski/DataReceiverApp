#include <Communication/http/IhttpHandler.h>
#include <Communication/http/httpHandler.h>
#include <Communication/http/httpHandlerMock.h>
#include <Communication/mqtt/mqttHandler.h>
#include <Communication/wifi/wiFiHandler.h>
#include <FS.h>
#include <HTTPClient.h>
#include <Json/jsonModels.h>
#include <Settings/ConfigManager.h>
#include <Settings/utils.h>
#include <display.h>
#include <freertos/FreeRTOS.h>
#include <freertos/queue.h>
#include <freertos/task.h>
#include <memory>
#include <string>

class DeviceMenager
{
public:
	static DeviceMenager &Instance()
	{
		static DeviceMenager instance;
		return instance;
	}

	DeviceMenager(const DeviceMenager &) = delete;
	DeviceMenager &operator=(const DeviceMenager &) = delete;
	DeviceMenager(DeviceMenager &&) = delete;
	DeviceMenager &operator=(DeviceMenager &&) = delete;

	void Run();
	void Stop();

private:
	DeviceMenager();
	~DeviceMenager() = default;

	static void mqttTask(void *pvParameters);
	static void wifiTask(void *pvParameters);
	static void setupTask(void *pvParameters);
	static void displayTask(void *pvParameters);

	void setupTaskImpl();
	void displayTaskImpl();
	void wifiTaskImpl();
	void mqttTaskImpl();

	QueueHandle_t wifiStatusQueue;

	TaskHandle_t setupTaskHandle = NULL;
	TaskHandle_t mqttTaskHandle = NULL;

	std::unique_ptr<IHttpHandler> httpHandler;
	HttpHandlerMock httptest;
	WiFiHandler wifiHandler;
	MQTTClientHandler mqttHandler;
	OLEDDisplay displayHandler;

	Json::Model::MqttConfig::MqttConfig mqttConfig;
};
