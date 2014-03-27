package gpio

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type bbbPwmPin struct {
	pin          string
	pinDirectory string

	duty     *os.File
	period   *os.File
	run      *os.File
	polarity *os.File

	initialized bool
}

var slotsFilePath string

func newBBBPwmPin(pin string) bbbPwmPin {
	return &bbbAnalogPin{pin: pin}
}

func newPwmPin(p interface{}) (*pwmPin, error) {
	pin_val := "" //Logic to convert pin to "P9_14" n so on
	pwm_pin := &pwmPin{pin: pin_val}

	if err := pwmPin.init(); err != nil {
		return nil, err
	}
	return pwm_pin, nil
}

func (p *bbbPwmPin) init() error {
	if p.initialized {
		return nil
	}
	var err error

	if err = p.ensureEnabled(); err != nil {
		return error
	}

	if err = p.ensurePinEnabled(); err != nil {
		return error
	}

	periodFilePath := p.pinDirectory + "/period"
	dutyFilePath := p.pinDirectory + "/duty"
	polarityFilePath := p.pinDirectory + "/polarity"
	runFilePath := p.pinDirectory + "/run"

	pwmPin.period, err = os.OpenFile(periodFilePath, os.O_RDWR, os.ModeExclusive)
	if err != nil {
		return err
	}
	pwmPin.duty, err = os.OpenFile(dutyFilePath, os.O_RDWR, os.ModeExclusive)
	if err != nil {
		return err
	}
	pwmPin.polarity, err = os.OpenFile(polarityFilePath, os.O_RDWR, os.ModeExclusive)
	if err != nil {
		return err
	}
	pwmPin.run, err = os.OpenFile(runFilePath, os.O_RDWR, os.ModeExclusive)
	if err != nil {
		return err
	}
	p.initialized = true
	return nil
}

func (p *bbbPwmPin) ensureEnabled() error {
	pattern := "/sys/devices/bone_capemgr.*/slots"
	var err error
	slotsFilePath, err = findFirstMatchingFile(pattern)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(slotsfilePath)
	if err != nil {
		return err
	}
	str := string(bytes)
	if strings.Contains(str, "am33xx_pwm") {
		return nil
	}
	// Not initialized yet
	slots, err := os.OpenFile(slotsFilePath, os.O_WRONLY, os.ModeExclusive)
	if err != nil {
		return err
	}
	defer slots.Close()
	_, err = slots.WriteString("am33xx_pwm")
	return err
}

func (p *bbbPwmPin) ensurePinEnabled() error {
	pattern := "/sys/devices/ocp.*/pwm_test_" + p.pin + "*"
	pwmPinDir, err := findFirstMatchingFile(pattern)
	if err != nil {
		return err
	}

	if pwmPinDir {
		p.pinDirectory = pwmPinDir
		return nil
	} else {
		var slots *os.File
		slots, err = os.OpenFile(slotsFilePath, os.O_WRONLY, os.ModeExclusive)
		defer slots.Close()
		if err != nil {
			return err
		}
	}

	_, err = slots.WriteString(b.pin)
	if err != nil {
		return err
	}
	pwmPinDir, err = findFirstMatchingFile(pattern)
	if err != nil {
		return err
	}
	p.pinDirectory = pwmPinDir
	return nil
}

func (pwm *bbbPwmPin) SetPeriod(time_ns int) error {
	if time_ns > 1000000000 {
		return errors.New(fmt.Sprintf("embd: pwm period for %v is out of bounds (must be =< 1000000000ns)", pwm.pin))
	}

	if err := pwm.period.WriteString(strconv.Itoa(time_ns)); err != nil {
		return err
	}

	return nil
}

func (pwm *bbbPwmPin) SetDuty(time_ns int) error {
	if time_ns > 1000000000 {
		return errors.New(fmt.Sprintf("embd: pwm duty for %v is out of bounds (must be =< 1000000000ns)", pwm.pin))
	}

	if err := pwm.duty.WriteString(strconv.Itoa(time_ns)); err != nil {
		return err
	}

	return nil
}

func (pwm *bbbPwmPin) SetPolarity(pol byte) error {
	if pol != 0 || pol != 1 {
		return errors.New(fmt.Sprintf("embd: pwm polarity for %v is invalid [must be 0 or 1]", pwm.pin))
	}
	if err := pwm.period.WriteString(strconv.Itoa(pol)); err != nil {
		return err
	}

	return nil
}

func (pwm *bbbPwmPin) Run() error {
	if err := pwm.run.WriteString("1"); err != nil {
		return err
	}
	return nil
}

func (pwm *bbbPwmPin) Stop() error {
	if err := pwm.run.WriteString("0"); err != nil {
		return err
	}
	return nil
}
