#pragma once
#include "IHttpHandler.h"

#include <string>

class HttpHandlerMock : public IHttpHandler
{
public:
	~HttpHandlerMock() override{};
	bool sendData(const std::string &destination, const std::string &jsonPayload) override
	{
		return true;
	}

	bool sendData(const std::string &destination, const std::string &jsonPayload, std::string &response) override
	{
		response = R"({
                    "MQQT_SERVER_IP" : "74.248.137.150",
					"MQQT_PORT" : 1883,
					"MQQT_USER" : "mqqtUser",
					"MQQT_PASSWORD" : "mqttPass",
					"MQQT_DATA_TOPIC" : "Device/38/Measurments",
					"MQQT_SERVICE_TOPIC" : "Device/38/Service",
					"MQQT_FIRMWARE_TOPIC" : "Device/38/Firmware"})";
		return true;
	}
};
