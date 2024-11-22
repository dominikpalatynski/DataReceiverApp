#ifndef JSON_GENERATOR_H
#define JSON_GENERATOR_H

#include <ArduinoJson.h>
#include <String>
#include <ctime>
#include <map>

class JsonBuilder
{
public:
	JsonBuilder();
	virtual ~JsonBuilder();

protected:
	template <typename DataType>
	void AddProperty(const String &name, const DataType &value);
	void AddPropertyTable(const String &name);
	template <typename DataType>
	void AddPropertyTableValue(const String &name, const String &dataName, const DataType &value);
	String generateJson();

private:
	JsonDocument doc;
	std::map<String, JsonObject> propertyTable;
};

class JsonSensorDataGenerator : JsonBuilder
{
public:
	JsonSensorDataGenerator();
	void initializeJson(const String &deviceID);
	void addSensorData(const float data, const String &sensorID);
	String generateJsonString();

private:
	String getTimeStamp();
};

class JsonServerRequestID : JsonBuilder
{
public:
	JsonServerRequestID();
	void initializeJson(const String &macAdress, const String &token);
	String generateJsonString();
};

#endif // JSON_GENERATOR_H
