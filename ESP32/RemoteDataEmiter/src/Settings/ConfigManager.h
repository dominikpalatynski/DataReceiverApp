#include "../Json/jsonModels.h"

#include <Preferences.h>
#include <string>
#include <type_traits>

template <typename T>
class ConfigManager
{
private:
	Preferences preferences;
	const std::string preferencesNamespace = "config";

	static_assert(std::is_base_of<JsonSerializable<T>, T>::value, "T must inherit from JsonSerializable.");

public:
	ConfigManager() = default;

	~ConfigManager()
	{
		preferences.end();
	}

	void save(const std::string &key, const T &data)
	{
		preferences.begin(preferencesNamespace.c_str(), false);
		std::string jsonData = data.toJsonString();
		preferences.putString(key.c_str(), jsonData.c_str());
		preferences.end();
	}

	bool load(const std::string &key, T &data)
	{
		preferences.begin(preferencesNamespace.c_str(), true);
		if(!preferences.isKey(key.c_str()))
		{
			preferences.end();
			return false;
		}

		std::string jsonString = preferences.getString(key.c_str(), "").c_str();
		preferences.end();

		try
		{
			data.fromJsonString(jsonString);
			return true;
		}
		catch(const std::exception &e)
		{
			Serial.println("Error during deserialization: ");
			Serial.println(e.what());
			return false;
		}
	}

	void remove(const std::string &key)
	{
		preferences.begin(preferencesNamespace.c_str(), false);
		preferences.remove(key.c_str());
		preferences.end();
	}
};
