#include "JsonGenerator.h"

JsonBuilder::JsonBuilder() : doc()
{
	doc.clear();
}

JsonBuilder::~JsonBuilder()
{
}

template <typename DataType>
void JsonBuilder::AddProperty(const String &name, const DataType &value)
{
	doc[name] = value;
}
void JsonBuilder::AddPropertyTable(const String &name)
{
	doc[name] = JsonArray();
	JsonObject data = doc[name].add<JsonObject>();
	propertyTable.insert({name, data});
}

template <typename DataType>
void JsonBuilder::AddPropertyTableValue(const String &name, const String &dataName, const DataType &value)
{
	JsonObject &data = propertyTable.at(name);
	data[dataName] = value;
}

String JsonBuilder::generateJson()
{
	String output;
	serializeJson(doc, output);
	return output;
}

JsonSensorDataGenerator::JsonSensorDataGenerator()
{
}

String JsonSensorDataGenerator::getTimeStamp()
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

void JsonSensorDataGenerator::initializeJson(const String &deviceID)
{
	const String timeStamp = getTimeStamp();
	const String timeStamp_str = "timeStamp";
	const String deviceID_str = "DeviceID";
	AddProperty(timeStamp_str, timeStamp);
	AddProperty(deviceID_str, deviceID);
}

void JsonSensorDataGenerator::addSensorData(const float data, const String &sensorID)
{
	const String sensors_str = "sensors";
	AddPropertyTable(sensors_str);
	AddPropertyTableValue("sensors", sensorID, data);
}

String JsonSensorDataGenerator::generateJsonString()
{
	return generateJson();
}

JsonServerRequestID::JsonServerRequestID()
{
}

void JsonServerRequestID::initializeJson(const String &macAdress, const String &token)
{
	const String mac_str = "mac";
	const String token_str = "token";
	AddProperty(mac_str, macAdress);
	AddProperty(token_str, token);
}

String JsonServerRequestID::generateJsonString()
{
	return generateJson();
}