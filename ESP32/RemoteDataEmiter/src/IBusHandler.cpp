#include "IBusHandler.h"

IBusHandler::IBusHandler(HardwareSerial &serialPort) : serial(&serialPort)
{
}

void IBusHandler::begin()
{
	serial->begin(115200, SERIAL_8N1, RX_PIN, -1);
	ibus.begin(*serial);
}

void IBusHandler::update()
{
	ibus.loop();
}

uint16_t IBusHandler::getChannelValue(int channel)
{
	if(channel >= 0 && channel < 14)
	{
		return ibus.readChannel(channel);
	}
	else
	{
		return 0;
	}
}
