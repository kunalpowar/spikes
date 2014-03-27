package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type PwmPin interface {
	SetPeriod(int) error
	SetDuty(int) error
	SetPolarity(int) error
	Run() error
	Stop() error
	Reset() error
}

type bbbPWMPin struct {
	pin          string
	pinDirectory string

	duty     *os.File
	period   *os.File
	run      *os.File
	polarity *os.File

	initialized bool
}

var slotsFilePath string

func newBBBPWMPin(pin string) PwmPin {
	return &bbbPWMPin{pin: pin}
}

func (p *bbbPWMPin) init() error {
	if p.initialized {
		return nil
	}
	var err error

	if err = p.ensureEnabled(); err != nil {
		return err
	}

	if err = p.ensurePinEnabled(); err != nil {
		return err
	}

	periodFilePath := p.pinDirectory + "/period"
	dutyFilePath := p.pinDirectory + "/duty"
	polarityFilePath := p.pinDirectory + "/polarity"
	runFilePath := p.pinDirectory + "/run"

	p.period, err = os.OpenFile(periodFilePath, os.O_RDWR, os.ModeExclusive)
	if err != nil {
		return err
	}
	p.duty, err = os.OpenFile(dutyFilePath, os.O_RDWR, os.ModeExclusive)
	if err != nil {
		return err
	}
	p.polarity, err = os.OpenFile(polarityFilePath, os.O_RDWR, os.ModeExclusive)
	if err != nil {
		return err
	}
	p.run, err = os.OpenFile(runFilePath, os.O_RDWR, os.ModeExclusive)
	if err != nil {
		return err
	}
	p.initialized = true
	return nil
}

func (p *bbbPWMPin) ensureEnabled() error {

	pattern := "/sys/devices/bone_capemgr.*/slots"
	slotsPath, err := findFirstMatchingFile(pattern)
	if err != nil {
		return err
	}
	slotsFilePath = slotsPath

	bytes, err := ioutil.ReadFile(slotsFilePath)
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

func (p *bbbPWMPin) ensurePinEnabled() error {

	pattern := "/sys/devices/ocp.*/pwm_test_" + p.pin + "*"
	pwmPinDir, err := findFirstMatchingFile(pattern)
	if err != nil {
		return err
	}

	var slots *os.File
	if pwmPinDir != "" {
		p.pinDirectory = pwmPinDir
		return nil
	} else {

		slots, err = os.OpenFile(slotsFilePath, os.O_WRONLY, os.ModeExclusive)
		defer slots.Close()
		if err != nil {
			return err
		}
	}

	pinBoneId := "bone_pwm_" + p.pin

	_, err = slots.WriteString(pinBoneId)
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

func (pwm *bbbPWMPin) SetPeriod(time_ns int) error {
	if err := pwm.init(); err != nil {
		return err
	}
	if time_ns > 1000000000 {
		return errors.New(fmt.Sprintf("embd: pwm period for %v is out of bounds (must be =< 1000000000ns)", pwm.pin))
	}

	if _, err := pwm.period.WriteString(strconv.Itoa(time_ns)); err != nil {
		return err
	}

	return nil
}

func (pwm *bbbPWMPin) SetDuty(time_ns int) error {
	if err := pwm.init(); err != nil {
		return err
	}

	if time_ns > 1000000000 {
		return errors.New(fmt.Sprintf("embd: pwm duty for %v is out of bounds (must be =< 1000000000ns)", pwm.pin))
	}

	if _, err := pwm.duty.WriteString(strconv.Itoa(time_ns)); err != nil {
		return err
	}

	return nil
}

func (pwm *bbbPWMPin) SetPolarity(pol int) error {
	if err := pwm.init(); err != nil {
		return err
	}

	if pol != 0 && pol != 1 {
		return errors.New(fmt.Sprintf("embd: pwm polarity for %v is invalid [must be 0 or 1]", pwm.pin))
	}
	if _, err := pwm.polarity.WriteString(strconv.Itoa(pol)); err != nil {
		return err
	}

	return nil
}

func (pwm *bbbPWMPin) Run() error {
	if err := pwm.init(); err != nil {
		return err
	}

	if _, err := pwm.run.WriteString("1"); err != nil {
		return err
	}
	return nil
}

func (pwm *bbbPWMPin) Stop() error {
	if err := pwm.init(); err != nil {
		return err
	}

	if _, err := pwm.run.WriteString("0"); err != nil {
		return err
	}
	return nil
}

func (pwm *bbbPWMPin) Reset() error {
	if err := pwm.SetPolarity(0); err != nil {
		return err
	}
	if err := pwm.SetDuty(0); err != nil {
		return err
	}
	if err := pwm.SetPeriod(1000000000); err != nil {
		return err
	}
	if err := pwm.Stop(); err != nil {
		return err
	}
	return nil
}

func findFirstMatchingFile(glob string) (string, error) {
	matches, err := filepath.Glob(glob)
	if err != nil {
		return "", err
	}
	if len(matches) >= 1 {
		return matches[0], nil
	}
	return "", nil
}

func main() {

	p := newBBBPWMPin("P9_14")
	p.Stop()
	p.Reset()
	defer p.Stop()

	err := p.SetPeriod(1000000)
	if err != nil {
		panic(err)
	}

	err = p.SetPolarity(0)
	if err != nil {
		panic(err)
	}
	p.Run()

	for i := 0; i < 5; i++ {
		err = p.SetDuty(500000 - (i * 100000))
		if err != nil {
			panic(err)
		}
		time.Sleep(1 * time.Second)
	}
	p.Stop()
}
