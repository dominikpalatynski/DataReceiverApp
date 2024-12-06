#pragma once
#include <string>

class IHttpHandler
{
public:
	virtual ~IHttpHandler() = default;

	virtual bool sendData(const std::string &destination, const std::string &jsonPayload) = 0;

	virtual bool sendData(const std::string &destination, const std::string &jsonPayload, std::string &response) = 0;
};
