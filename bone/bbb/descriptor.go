package bbb

import (
	"github.com/kidoman/embd/hal"
	"github.com/kidoman/embd/host/generic/linux/gpio"
	"github.com/kidoman/embd/host/generic/linux/i2c"
)

type descriptor struct {
	rev int
}

func (d *descriptor) GPIO() hal.GPIO {
	return gpio.New(Pins)
}

func (d *descriptor) I2C() hal.I2C {
	return i2c.New()
}

func Descriptor(rev int) *descriptor {
	return &descriptor{rev}
}
