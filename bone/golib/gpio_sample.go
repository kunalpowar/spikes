package main

import (
	"fmt"
	"gpio"
	"time"
)

func main() {
	io = gpio.Init()

	io.SetPinDir(29, gpio.Out)

	for {
		time.Sleep(1 * time.Second)
		if err = io.SetPinValue(29, io.High); err != nil {
			fmt.Errorf("Gpio: Could not assign value to pin %i due to %s", 29, err.Error())
		}
		time.Sleep(1 * time.Second)
		if err = io.SetPinValue(29, io.Low); err != nil {
			fmt.Errorf("Gpio: Could not assign value to pin %i due to %s", 29, err.Error())
		}
	}
}
