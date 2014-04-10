package embd

import (
	"os"
)

type Host interface {
	NewDigitalPin() (err error)
	NewAnalogPin() (err error)
}

type bbb struct {
}

func (b *bbb) NewDigitalPin() (err error) {
	return nil
}

func (b *bbb) NewAnalogPin() (err error) {
	return nil
}

func main() {

}
