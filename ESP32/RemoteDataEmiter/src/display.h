#include <Adafruit_GFX.h>
#include <Adafruit_SSD1306.h>

#define SCREEN_WIDTH 128
#define SCREEN_HEIGHT 64
#define OLED_RESET -1

#define SCL_PIN 18
#define SDA_PIN 19

class OLEDDisplay
{
private:
	Adafruit_SSD1306 display;
	String rows[6];

public:
	OLEDDisplay();
	void begin();
	void setRow(int row, String text);
	void clear();
	void updateDisplay();
};