#pragma once
#include "../jsonSerializerDeserializer.h"

namespace Json::Model::SensorsData
{
	struct Sensor
	{
		std::string data;
		std::string sensorID;

		NLOHMANN_DEFINE_TYPE_INTRUSIVE(Sensor, data, sensorID);
	};

	struct DeviceData : JsonSerializable<DeviceData>
	{
		std::string timeStamp;
		std::vector<Sensor> sensors;

		NLOHMANN_DEFINE_TYPE_INTRUSIVE(DeviceData, timeStamp, sensors);
	};

} // namespace Json::Model::SensorsData
