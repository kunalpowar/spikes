void setup()
{
	Serial.begin(9600);
	pinMode(5, INPUT);	
}

void loop()
{
	int val = analogRead(5);
	Serial.println(val);
}