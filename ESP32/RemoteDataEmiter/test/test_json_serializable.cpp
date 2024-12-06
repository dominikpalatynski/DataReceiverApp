#include "../src/Json/jsonModels.h"
#include "../src/Settings/utils.h"

#include <gtest/gtest.h>

using namespace Json::Model;

// Test: Serializacja i deserializacja MQTT Config
TEST(JsonSerializableTest, MqttConfigSerialization)
{
	MqttConfig::MqttConfig config("192.168.1.1", 1883, "user", "password", "data/topic", "service/topic",
								  "firmware/topic");
	std::string json = config.toJsonString();

	// Sprawdzenie czy dane są poprawnie serializowane
	EXPECT_NE(json.find("192.168.1.1"), std::string::npos);
	EXPECT_NE(json.find("1883"), std::string::npos);
	EXPECT_NE(json.find("user"), std::string::npos);

	// Deserializacja
	MqttConfig::MqttConfig deserializedConfig;
	deserializedConfig.fromJsonString(json);

	// Sprawdzenie poprawności deserializacji
	EXPECT_EQ(deserializedConfig.serverIp, "192.168.1.1");
	EXPECT_EQ(deserializedConfig.port, 1883);
	EXPECT_EQ(deserializedConfig.user, "user");
	EXPECT_EQ(deserializedConfig.password, "password");
	EXPECT_EQ(deserializedConfig.dataTopic, "data/topic");
	EXPECT_EQ(deserializedConfig.serviceTopic, "service/topic");
	EXPECT_EQ(deserializedConfig.firmwareTopic, "firmware/topic");
}

// Test: Serializacja i deserializacja SensorConfig
TEST(JsonSerializableTest, SensorsConfigSerialization)
{
	SensorConfig::Wires wires(1, 2, 3, 4, 5, 6);
	SensorConfig::Settings settings("UART", 9600, "NONE", 0, 0, "", wires);
	SensorConfig::Sensor sensor("sensor_1", settings);

	SensorConfig::SensorsConfig config("SensorType", {sensor});
	std::string json = config.toJsonString();

	// Sprawdzenie czy dane są poprawnie serializowane
	EXPECT_NE(json.find("sensor_1"), std::string::npos);
	EXPECT_NE(json.find("UART"), std::string::npos);
	EXPECT_NE(json.find("NONE"), std::string::npos);

	// Deserializacja
	SensorConfig::SensorsConfig deserializedConfig;
	deserializedConfig.fromJsonString(json);

	// Sprawdzenie poprawności deserializacji
	EXPECT_EQ(deserializedConfig.configType, "SensorType");
	ASSERT_EQ(deserializedConfig.sensors.size(), 1);
	EXPECT_EQ(deserializedConfig.sensors[0].sensorID, "sensor_1");
	EXPECT_EQ(deserializedConfig.sensors[0].settings.protocol, "UART");
	EXPECT_EQ(deserializedConfig.sensors[0].settings.baudRate, 9600);
	EXPECT_EQ(deserializedConfig.sensors[0].settings.parity, "NONE");
}

// Test: Serializacja i deserializacja SensorsData
TEST(JsonSerializableTest, SensorsDataSerialization)
{
	SensorsData::Sensor sensor("temperature:25.5", "sensor_1");
	SensorsData::DeviceData deviceData("2023-12-04T10:00:00Z", "device_123", {sensor});
	std::string json = deviceData.toJsonString();

	// Sprawdzenie czy dane są poprawnie serializowane
	EXPECT_NE(json.find("temperature:25.5"), std::string::npos);
	EXPECT_NE(json.find("sensor_1"), std::string::npos);
	EXPECT_NE(json.find("2023-12-04T10:00:00Z"), std::string::npos);

	// Deserializacja
	SensorsData::DeviceData deserializedData;
	deserializedData.fromJsonString(json);

	// Sprawdzenie poprawności deserializacji
	EXPECT_EQ(deserializedData.timeStamp, "2023-12-04T10:00:00Z");
	EXPECT_EQ(deserializedData.deviceID, "device_123");
	ASSERT_EQ(deserializedData.sensors.size(), 1);
	EXPECT_EQ(deserializedData.sensors[0].data, "temperature:25.5");
	EXPECT_EQ(deserializedData.sensors[0].sensorID, "sensor_1");
}

void setup()
{
	// should be the same value as for the `test_speed` option in "platformio.ini"
	// default value is test_speed=115200
	Serial.begin(115200);

	::testing::InitGoogleTest();
}

void loop()
{
	// Run tests
	if(RUN_ALL_TESTS())
		;

	// sleep 1 sec
	delay(1000);
}