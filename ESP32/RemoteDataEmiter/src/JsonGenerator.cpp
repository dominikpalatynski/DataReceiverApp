#include "JsonGenerator.h"

JsonGenerator::JsonGenerator() : doc()
{
}

String JsonGenerator::getTimeStamp()
{
	time_t now;
	time(&now);

	now += 3600;

	struct tm timeinfo;
	gmtime_r(&now, &timeinfo);

	char buffer[30];

	strftime(buffer, sizeof(buffer), "%Y-%m-%dT%H:%M:%S+01:00", &timeinfo);

	return String(buffer);
}

void JsonGenerator::initializeJson(const String deviceID)
{
	doc.clear();
	doc["timeStamp"] = getTimeStamp();
	doc["deviceID"] = deviceID;
	doc["sensors"] = JsonArray();
}

void JsonGenerator::addSensorData(const float data, const String sensorID)
{
	JsonObject sensorData = doc["sensors"].add<JsonObject>();
	sensorData["data"] = data;
	sensorData["sensorID"] = sensorID;
}

String JsonGenerator::generateJsonString()
{
	String output;
	serializeJson(doc, output);
	return output;
}
