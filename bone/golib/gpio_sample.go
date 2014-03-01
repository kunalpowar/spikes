package main

func main() {

	//Gpio Output
	led = gpio.NewPin(31, gpio.Out)

	led.DigitalWrite(gpio.High)
	led.setHigh()

	//Gpio Input
	switchIp = gpio.NewPin(29, gpio.In)
	if switchIp.DigitalRead() == gpio.Low {
		led.SetLow()
	} else {
		led.SetHigh()
	}

	//Pwm
	pwmChannel = gpio.NewPwmChannel(21)

	pwmChannel.setPeriod(10000)
	pwmChannel.setDutyCycle(5000)

	pwmChannel.start()
	pwmChannel.setDutyCycle(500)
	pwmChannel.stop()
}
