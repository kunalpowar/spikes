package gpio

import (
	"os"
)

type pwm struct {
	exporter   *os.File
	unexporter *os.File

	pin *pwmPin
}

type pwmPin struct {
	duty_ns   *os.File
	period_ns *os.File

	polarity *os.File
}

func newPwmPin(num int) (*pwm, error) {
	pwm := &pwm{}
}

func (pwm *pwm) init() error {
	file, err := os.OpenFile("/sys/devices/bone_capemgr.8/slots", os.O_RDWR, os.ModeExclusive)
	if err != nil {
		return err
	}

	if !pwmInitialized() {
		_, err = file.Write("am33xx_pwm")
		if err != nil {
			return err
		}
	}

	pwm.exporter, err = os.OpenFile("/sys/class/pwm/export", os.O_WRONLY, os.ModeExclusive)
	if err != nil {
		return err
	}

	pwm.unexporter, err = os.OpenFile("/sys/class/pwm/unexport", os.O_WRONLY, os.ModeExclusive)
	if err != nil {
		return err
	}
}
