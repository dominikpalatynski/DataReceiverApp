#pragma once
#include "../jsonSerializerDeserializer.h"

namespace Json::Model::WifiConfig
{
	struct WifiConfig : JsonSerializable<WifiConfig>
	{
		std::string ssid;
		std::string password;

		NLOHMANN_DEFINE_TYPE_INTRUSIVE(WifiConfig, ssid, password);
	};
} // namespace Json::Model::WifiConfig
