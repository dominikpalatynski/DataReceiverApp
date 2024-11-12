#ifndef JSON_GENERATOR_H
#define JSON_GENERATOR_H

#include <ArduinoJson.h>
#include <String>
#include <ctime>

class JsonGenerator
{
public:
	JsonGenerator();
	void initializeJson(const String deviceID);
	void addSensorData(const float data, const String sensorID);
	String generateJsonString();

private:
	JsonDocument doc;

	String getTimeStamp();
};

#endif // JSON_GENERATOR_H
