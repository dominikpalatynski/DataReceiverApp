#pragma once
#include "../jsonSerializerDeserializer.h"

namespace Json::Model::SensorConfig
{
	struct Wires
	{
		int RX = -1;
		int TX = -1;
		int SCL = -1;
		int SDA = -1;
		int IN = -1;
		int GND = -1;

		NLOHMANN_DEFINE_TYPE_INTRUSIVE(Wires, RX, TX, SCL, SDA, IN, GND);
	};

	struct Settings
	{
		std::string protocol;
		int baudRate = 0;	// Only for UART
		std::string parity; // Only for UART
		int address = 0;	// Only for I2C
		int frequency = 0;	// Only for I2C
		std::string type;	// Only for Analog
		Wires wires;		// Common for all protocols

		NLOHMANN_DEFINE_TYPE_INTRUSIVE(Settings, protocol, baudRate, parity, address, frequency, type, wires);
	};

	struct Sensor
	{
		std::string sensorID;
		Settings settings;

		NLOHMANN_DEFINE_TYPE_INTRUSIVE(Sensor, sensorID, settings);
	};

	struct SensorsConfig : JsonSerializable<SensorsConfig>
	{
		std::string configType;
		std::vector<Sensor> sensors;

		NLOHMANN_DEFINE_TYPE_INTRUSIVE(SensorsConfig, configType, sensors);
	};

} // namespace Json::Model::SensorConfig
