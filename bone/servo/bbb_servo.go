package main

import (
	"fmt"
	"github.com/kidoman/embd"
	"time"
)

type Servo interface {
	SetAngle(int) error
	Reset() error
}


type servo struct {
	pin       string
	pwmDriver embd.PWMPin

	initialized bool
}

func newServo(pin string) Servo {
	return &servo{pin: pin}
}

func (s *servo) init() error {
	if s.initialized {
		return nil
	}

	var p embd.PWMPin
	var err error

	fmt.Printf("NewPWMPin called for: %v", s.pin)
	p, err = embd.NewPWMPin(s.pin)
	if err != nil {
		return err
	}
	s.pwmDriver = p
	fmt.Printf("Setting polarity here to: %v", embd.Positive)
	if err = s.pwmDriver.SetPolarity(embd.Positive); err != nil {
		return err
	}

	fmt.Printf("Setting period here to: %v", 15000000)
	if err = s.pwmDriver.SetPeriod(15000000); err != nil {
		return err
	}

	fmt.Printf("Setting duty here to: %v", 2500000)
	if err = s.pwmDriver.SetDuty(2500000); err != nil {
		return err
	}
	s.initialized = true
	return nil
}

func (s *servo) SetAngle(angle int) error {
	if err := s.init(); err != nil {
		return err
	}
	duty := mymap(angle, 0, 180, 1000000, 2500000)

	fmt.Printf("Setting duty from server to: %v\n", duty)
	if err := s.pwmDriver.SetDuty(duty); err != nil {
		return err
	}
	return nil
}

func (s *servo) Reset() error {
	if err := s.init(); err != nil {
		return err
	}
	duty := mymap(90, 0, 180, 1000000, 2500000)

	if err := s.pwmDriver.SetDuty(duty); err != nil {
		return err
	}
	return nil
}

func mymap(x, inmin, inmax, outmin, outmax int) int {
	return (x-inmin)*(outmax-outmin)/(inmax-inmin) + outmin
}

func main() {
	fmt.Println("initGpio called")
	var err error
	if err = embd.InitGPIO(); err != nil {
		panic(err)
	}
	defer embd.CloseGPIO()
	s := newServo("P9_14")

	for i := 0; i < 18; i++ {
		angle := 10 + (10 * i)
		fmt.Printf("Setting angle: %v\n", angle)
		if err := s.SetAngle(angle); err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
}
