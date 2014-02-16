package main

import (
	"gpio"
	"time"
)

func main() {
	io = gpio.Init()

	io.SetPinDir(29, gpio.Out)

	for {
		time.Sleep(1 * time.Second)
		io.SetPinValue(29, io.High)

		time.Sleep(1 * time.Second)
		io.SetPinValue(29, io.Low)
	}
}
