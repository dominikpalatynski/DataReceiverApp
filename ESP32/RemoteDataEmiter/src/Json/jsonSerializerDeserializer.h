#pragma once
#include <nlohmann/json.hpp>
#include <string>
#include <vector>

template <typename T>
class JsonSerializable
{
public:
	virtual ~JsonSerializable() = default;

	std::string toJsonString() const
	{
		try
		{
			nlohmann::json j = static_cast<const T &>(*this);
			return j.dump();
		}
		catch(const std::exception &e)
		{
			Serial.println("Error during JSON serialization:");
			Serial.println(e.what());
			return "{}";
		}
	}

	void fromJsonString(const std::string &jsonString)
	{
		try
		{
			nlohmann::json j = nlohmann::json::parse(jsonString);
			static_cast<T &>(*this) = j.get<T>();
		}
		catch(const nlohmann::json::parse_error &e)
		{
			Serial.println("JSON parsing error:");
			Serial.println(e.what());
		}
		catch(const nlohmann::json::type_error &e)
		{
			Serial.println("JSON type error:");
			Serial.println(e.what());
		}
		catch(const std::exception &e)
		{
			Serial.println("Error during JSON deserialization:");
			Serial.println(e.what());
		}
		catch(...)
		{
			Serial.println("Unknown error during JSON deserialization");
		}
	}
};