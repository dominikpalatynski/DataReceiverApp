#ifndef IBusHandler_h
#define IBusHandler_h

#include <Arduino.h>
#include <IBusBM.h>

#define RX_PIN 10

class IBusHandler
{
private:
	IBusBM ibus;
	HardwareSerial *serial;

public:
	IBusHandler(HardwareSerial &serialPort);

	void begin();

	void update();

	uint16_t getChannelValue(int channel);
};

#endif
