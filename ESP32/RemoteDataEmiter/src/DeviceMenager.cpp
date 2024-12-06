#include "DeviceMenager.h"

void DeviceMenager::setupTask(void *pvParameters)
{
	DeviceMenager *instance = static_cast<DeviceMenager *>(pvParameters);
	instance->setupTaskImpl();
}

void DeviceMenager::displayTask(void *pvParameters)
{
	DeviceMenager *instance = static_cast<DeviceMenager *>(pvParameters);
	instance->displayTaskImpl();
}

void DeviceMenager::wifiTask(void *pvParameters)
{
	DeviceMenager *instance = static_cast<DeviceMenager *>(pvParameters);
	instance->wifiTaskImpl();
}

void DeviceMenager::mqttTask(void *pvParameters)
{
	DeviceMenager *instance = static_cast<DeviceMenager *>(pvParameters);
	instance->mqttTaskImpl();
}

void DeviceMenager::Run()
{
	xTaskCreate(DeviceMenager::setupTask, "Setup Task", 4096, NULL, 1, &setupTaskHandle);
	xTaskCreate(DeviceMenager::displayTask, "Display Task", 4096, NULL, 1, NULL);
	xTaskCreate(DeviceMenager::wifiTask, "WiFi Task", 4096, NULL, 1, NULL);
	//  xTaskCreate(mqttTask, "MQTT Task", 4096, NULL, 1, &mqttTaskHandle);
}

void DeviceMenager::Stop()
{
}

DeviceMenager::DeviceMenager()
{
	Serial.begin(115200);

	wifiStatusQueue = xQueueCreate(1, sizeof(bool));
	if(wifiStatusQueue == NULL)
	{
		Serial.println("Can't create wifiStatus queue");
		return;
	}

	bool useHttpMock = true;
	if(useHttpMock)
	{
		httpHandler = std::make_unique<HttpHandlerMock>();
	}
	else
	{
		httpHandler = std::make_unique<HttpHandler>();
	}
}

void DeviceMenager::mqttTaskImpl()
{
	while(true)
	{
		ulTaskNotifyTake(pdFALSE, portMAX_DELAY); // Wait for Wifi

		if(!mqttHandler.isConnected())
		{
			mqttHandler.begin(mqttConfig.MQQT_SERVER_IP, mqttConfig.MQQT_PORT, mqttConfig.MQQT_USER,
							  mqttConfig.MQQT_PASSWORD);
			// mqttHandler.setCallback(mqttCallback);
		}

		// publishData();

		vTaskDelay(5000 / portTICK_PERIOD_MS);
	}
}

void DeviceMenager::wifiTaskImpl()
{
	bool wifiConnected = false;
	while(true)
	{
		if(wifiHandler.isConnected())
		{
			wifiConnected = true;
			xQueueOverwrite(wifiStatusQueue, &wifiConnected);
			xTaskNotifyGive(mqttTaskHandle);
			xTaskNotifyGive(setupTaskHandle);
			vTaskDelay(5000 / portTICK_PERIOD_MS);
			continue;
		}
		wifiConnected = false;
		xQueueOverwrite(wifiStatusQueue, &wifiConnected);
		xTaskNotify(mqttTaskHandle, 0, eNotifyAction::eSetBits);
		wifiHandler.connect(Utils::wifi_ssid, Utils::wifi_password);
	}
}

void DeviceMenager::setupTaskImpl()
{
	bool forceConfigRequest = false;

	ConfigManager<Json::Model::MqttConfig::MqttConfig> mqttConfigManager;
	mqttConfigManager.remove(Utils::mqtt_config);
	while(true)
	{
		Serial.println("Loading mqtt config...");
		if(mqttConfigManager.load(Utils::mqtt_config, mqttConfig) && !forceConfigRequest)
		{
			Serial.println("Config loaded successfully!");
			vTaskDelete(NULL);
		}
		Serial.println("Config is empty, sending request...");

		// ulTaskNotifyTake(pdFALSE, portMAX_DELAY); // Wait for Wifi

		std::string mac = wifiHandler.getMacAdress();

		Json::Model::MqttConfigRequest::MqttConfigRequest deviceInfoRequest;
		deviceInfoRequest.mac = mac;
		deviceInfoRequest.token = Utils::http_token;

		std::string deviceInfoRequestJson = deviceInfoRequest.toJsonString();

		std::string responseJson;

		if(httptest.sendData(Utils::http_endpoint, deviceInfoRequestJson, responseJson))
		{
			Serial.println("Config downloaded successfully!");
			mqttConfig.fromJsonString(responseJson);

			std::string configStr = mqttConfig.toJsonString();
			Serial.println(configStr.c_str());

			Serial.println("Saving config...");
			mqttConfigManager.save(Utils::mqtt_config, mqttConfig);
			Serial.println("Saving complete");
			// vTaskDelete(NULL);
		}
	}
	Serial.println("Failed to download mqtt config! Retrying in 5 seconds...");
	vTaskDelay(5000 / portTICK_PERIOD_MS);
}

void DeviceMenager::displayTaskImpl()
{
	bool wifiConnected = false;

	displayHandler.begin();
	displayHandler.setRow(0, "Wifi:");
	displayHandler.setRow(1, "not connected");
	displayHandler.setRow(2, "MQTT topic:");
	displayHandler.setRow(4, "MQTT data:");

	while(true)
	{
		xQueuePeek(wifiStatusQueue, &wifiConnected, portMAX_DELAY);

		if(wifiConnected)
		{
			displayHandler.setRow(1, wifiHandler.getSSID().c_str());
		}
		else
		{
			displayHandler.setRow(1, "not connected");
		}

		vTaskDelay(1000 / portTICK_PERIOD_MS);
	}
}
