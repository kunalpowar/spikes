#include <Servo.h>

Servo servo;

void setup()
{
	Serial.begin(9600);
	pinMode(5, INPUT);	
	servo.attach(3);
}

void loop()
{
	int val = analogRead(5);

	int angle = map(val, 506, 680, 0, 180);
	servo.write(angle);
	Serial.println(val);
}