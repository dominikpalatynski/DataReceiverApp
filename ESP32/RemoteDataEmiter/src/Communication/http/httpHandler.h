#pragma once
#include "IHttpHandler.h"

#include <HTTPClient.h>
#include <string>

class HttpHandler : public IHttpHandler
{
private:
	HTTPClient httpClient;

public:
	~HttpHandler() override{};

	bool sendData(const std::string &destination, const std::string &jsonPayload) override
	{
		httpClient.begin(destination.c_str());
		httpClient.addHeader("Content-Type", "application/json");

		int httpResponseCode = httpClient.POST(jsonPayload.c_str());
		httpClient.end();

		return httpResponseCode > 0 && httpResponseCode < 300;
	}

	bool sendData(const std::string &destination, const std::string &jsonPayload, std::string &response) override
	{
		httpClient.begin(destination.c_str());
		httpClient.addHeader("Content-Type", "application/json");

		int httpResponseCode = httpClient.POST(jsonPayload.c_str());
		if(httpResponseCode > 0)
		{
			response = httpClient.getString().c_str();
		}
		httpClient.end();

		return httpResponseCode > 0 && httpResponseCode < 300;
	}
};
