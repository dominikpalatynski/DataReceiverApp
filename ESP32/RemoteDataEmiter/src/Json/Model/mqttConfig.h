#pragma once
#include "../jsonSerializerDeserializer.h"

namespace Json::Model::MqttConfig
{
	struct MqttConfig : JsonSerializable<MqttConfig>
	{
		std::string MQQT_SERVER_IP;
		int MQQT_PORT;
		std::string MQQT_USER;
		std::string MQQT_PASSWORD;
		std::string MQQT_DATA_TOPIC;
		std::string MQQT_SERVICE_TOPIC;
		std::string MQQT_FIRMWARE_TOPIC;

		NLOHMANN_DEFINE_TYPE_INTRUSIVE(MqttConfig, MQQT_SERVER_IP, MQQT_PORT, MQQT_USER, MQQT_PASSWORD, MQQT_DATA_TOPIC,
									   MQQT_SERVICE_TOPIC, MQQT_FIRMWARE_TOPIC);
	};
} // namespace Json::Model::MqttConfig
