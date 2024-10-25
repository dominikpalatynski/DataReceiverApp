#include "display.h"

OLEDDisplay::OLEDDisplay() : display(SCREEN_WIDTH, SCREEN_HEIGHT, &Wire, OLED_RESET)
{
}

void OLEDDisplay::begin()
{
	Wire.begin(SDA_PIN, SCL_PIN);
	if(!display.begin(SSD1306_SWITCHCAPVCC, 0x3C))
	{
		Serial.println(F("OLED display doesn't found!"));
		return;
	}
	display.clearDisplay();
	display.display();
	display.setTextColor(SSD1306_WHITE);
}

void OLEDDisplay::setRow(int row, String text)
{
	if(row >= 0 && row < 6)
	{
		rows[row] = text;
		updateDisplay();
	}
}

void OLEDDisplay::clear()
{
	for(int i = 0; i < 6; i++)
	{
		rows[i] = "";
	}
	updateDisplay();
}

void OLEDDisplay::updateDisplay()
{
	display.clearDisplay();
	display.setTextSize(1);

	for(int i = 0; i < 6; i++)
	{
		display.setCursor(0, (i * 10));
		display.println(rows[i]);
	}
	display.display();
}