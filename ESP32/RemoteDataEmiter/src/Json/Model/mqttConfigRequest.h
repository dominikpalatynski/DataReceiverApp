#pragma once
#include "../jsonSerializerDeserializer.h"

namespace Json::Model::MqttConfigRequest
{
	struct MqttConfigRequest : JsonSerializable<MqttConfigRequest>
	{
		std::string token;
		std::string mac;

		NLOHMANN_DEFINE_TYPE_INTRUSIVE(MqttConfigRequest, token, mac);
	};

} // namespace Json::Model::MqttConfigRequest
