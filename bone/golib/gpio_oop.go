package gpio

import (
	"errors"
	"os"
	"strconv"
)

const (
	HIGH = int(1)
	LOW  = int(0)

	OUT = "out"
	IN  = "in"
)

var pinExporter, _ = os.OpenFile("/sys/class/gpio/export", os.O_WRONLY, os.ModeExclusive)
var pinUnexporter, _ = os.OpenFile("/sys/class/gpio/unexport", os.O_WRONLY, os.ModeExclusive)

type Pin interface {
	SetDirection(direction string) (err error)
	SetHigh() (err error)
	SetLow() (err error)
	DigitalWrite() (err error)

	GetValue() (val int, err error)
	DigitalRead() (val int, err error)
	GetDirection() (direction string, err error)
}

type PwmChannel interface {
	AnalogWrite(analogValue uint32) (err error)
	SetDutyCycle(dutyCycle_ns uint32) (err error)
	SetPeriod(period_ns uint32) (err error)
	SetFrequency(freq uint32, err error)

	GetDutyCycle() (dutyCycle_ns uint32, err error)
	GetPeriod() (period_ns uint32, err error)
	GetFrequency() (freq uint32, err error)
}

type AnalogInput interface {
	AnalogRead()
}

type SpiChannel interface {
}

type I2cChannel interface {
}

type pin struct {
	pinExporter   *os.File
	pinUnexporter *os.File

	number int

	currentDir   string
	currentValue int

	dir *os.File
	val *os.File
}

func NewPin(num int) Pin {
	var p *pin
	p = &pin

	if p.pinExporter, err = os.OpenFile("/sys/class/gpio/export", os.O_WRONLY, os.ModeExclusive); err != nil {
		return errors.New(fmt.Sprintf("Gpio: cannot open export file due to %s", err.Error()))
	}
	if p.pinUnexporter, err = os.OpenFile("/sys/class/gpio/unexport", os.O_WRONLY, os.ModeExclusive); err != nil {
		return errors.New(fmt.Sprintf("Gpio: cannot open unexport file due to %s", err.Error()))
	}

	return p
}

func (p *pin) SetDirection(dir string) (err error) {
	if p.dir == nil {
		if _, err = PinExporter.WriteString(strconv.Itoa(p.number)); err != nil {
			return
		}

		pin_dir_file_path := fmt.Sprintf("/sys/class/gpio/gpio%i/direction", p.number)
		if p.dir, err = os.OpenFile(pin_dir_file_path, os.O_RDWR, os.ModeExclusive); err != nil {
			return
		}
	}
	if _, err = p.dir.WriteString(dir); err != nil {
		return
	}
	p.currentDir = dir
	return
}

func (p *pin) DigitalWrite(pinValue int) (err error) {
	if p.dir == nil {
		p.SetDirection(OUT)
	}

	if p.val == nil {
		pin_val_file_path := fmt.Sprintf("/sys/class/gpio/gpio%i/value", p.number)

		if p.val, err = os.OpenFile(pin_val_file_path, os.O_RDWR, os.ModeExclusive); err != nil {
			return
		}
	}

	if _, err = p.val.WriteString(strconv.Itoa(pinValue)); err != nil {
		return
	}
	p.currentValue = pinValue
	return
}

func (p *pin) SetHigh() (err error) {
	return p.DigitalWrite(HIGH)
}

func (p *pin) SetLow() (err error) {
	return p.DigitalWrite(LOW)
}

func (p *pin) DigitalRead() (value string, err error) {
	if p.dir == nil {
		return Low, errors.New(fmt.Sprintf("Gpio: direction not set for pin %i", p.number))
	}

	if p.val == nil {
		pin_val_file_path := fmt.Sprintf("/sys/class/gpio/gpio%i/value", p.number)

		p.val, err = os.OpenFile(pin_val_file_path, os.O_RDWR, os.ModeExclusive)
		if err != nil {
			return
		}
	}
	var num int
	var byteArray = make([]byte, 10)
	if num, err = p.val.Read(byteArray); err != nil {
		return
	}
	status := string(byteArray[:num])
	if status == strconv.Itoa(1) {

		return High, nil
	} else {
		return Low, nil
	}
}
