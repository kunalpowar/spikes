package gpio

import (
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

const (
	Out = "out"
	In  = "in"

	High = 1
	Low  = 0
)

var allowedPins = map[int]string{
	//GPIO_No_____|_Head_Pin_|_$PINS__|__________________________Notes________________________________
	30:  "allowed", // P9_11   |   28   |   GPIOs limit current to 4‐6mA output and approx. 8mA on input.
	60:  "allowed", // P9_12   |   30   |
	31:  "allowed", // P9_13   |   29   |
	50:  "allowed", // P9_14   |   18   |
	48:  "allowed", // P9_15   |   16   |
	51:  "allowed", // P9_16   |   19   |
	5:   "allowed", // P9_17   |   87   |
	4:   "allowed", // P9_18   |   86   |
	13:  "used",    // P9_19   |   95   |   Allocated (Group: pinmux_i2c2_pins)
	12:  "used",    // P9_20   |   94   |   Allocated (Group: pinmux_i2c2_pins)
	3:   "allowed", // P9_21   |   85   |
	2:   "allowed", // P9_22   |   84   |
	49:  "allowed", // P9_23   |   17   |
	15:  "allowed", // P9_24   |   97   |
	117: "used",    // P9_25   |   107  |   Allocated (Group: mcasp0_pins)
	14:  "allowed", // P9_26   |   96   |
	115: "allowed", // P9_27   |   105  |
	113: "used",    // P9_28   |   103  |   Allocated (Group: mcasp0_pins)
	111: "used",    // P9_29   |   101  |   Allocated (Group: mcasp0_pins)
	112: "allowed", // P9_30   |   102  |
	110: "used",    // P9_31   |   100  |   Allocated (Group: mcasp0_pins)

	38: "used",    //  P8_03   |   6    |   Used on Board (Group: pinmux_emmc2_pins)
	39: "used",    //  P8_04   |   7    |   Used on Board (Group: pinmux_emmc2_pins)
	34: "used",    //  P8_05   |   2    |   Used on Board (Group: pinmux_emmc2_pins)
	35: "used",    //  P8_06   |   3    |   Used on Board (Group: pinmux_emmc2_pins)
	66: "allowed", //  P8_07   |   36   |
	67: "allowed", //  P8_08   |   37   |
	69: "allowed", //  P8_09   |   39   |
	68: "allowed", //  P8_10   |   38   |
	45: "allowed", //  P8_11   |   13   |
	44: "allowed", //  P8_12   |   12   |
	23: "allowed", //  P8_13   |   9    |
	26: "allowed", //  P8_14   |   10   |
	47: "allowed", //  P8_15   |   15   |
	46: "allowed", //  P8_16   |   14   |
	27: "allowed", //  P8_17   |   11   |
	65: "allowed", //  P8_18   |   35   |
	22: "allowed", //  P8_19   |   8    |
	63: "used",    //  P8_20   |   33   |	  Used on Board (Group: pinmux_emmc2_pins)
	62: "used",    //  P8_21   |   32   |	  Used on Board (Group: pinmux_emmc2_pins)
	37: "used",    //  P8_22   |   5    |	  Used on Board (Group: pinmux_emmc2_pins)
	36: "used",    //  P8_23   |   4    |	  Used on Board (Group: pinmux_emmc2_pins)
	33: "used",    //  P8_24   |   1    |	  Used on Board (Group: pinmux_emmc2_pins)
	32: "used",    //  P8_25   |   0    |	  Used on Board (Group: pinmux_emmc2_pins)
	61: "allowed", //  P8_26   |   31   |
	86: "used",    //  P8_27   |   56   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	88: "used",    //  P8_28   |   58   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	87: "used",    //  P8_29   |   57   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	89: "used",    //  P8_30   |   59   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	10: "used",    //  P8_31   |   54   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	11: "used",    //  P8_32   |   55   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	9:  "used",    //  P8_33   |   53   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	81: "used",    //  P8_34   |   51   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	8:  "used",    //  P8_35   |   52   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	80: "used",    //  P8_36   |   50   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	78: "used",    //  P8_37   |   48   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	79: "used",    //  P8_38   |   49   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	76: "used",    //  P8_39   |   46   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	77: "used",    //  P8_40   |   47   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	74: "used",    //  P8_41   |   44   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	75: "used",    //  P8_42   |   45   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	72: "used",    //  P8_43   |   42   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	73: "used",    //  P8_44   |   43   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	70: "used",    //  P8_45   |   40   |   Allocated (Group: nxp_hdmi_bonelt_pins)
	71: "used",    //  P8_46   |   41   |   Allocated (Group: nxp_hdmi_bonelt_pins)
}

type Gpio interface {
	SetPinDir(pin int, pinDir string) (err error)

	SetPinValue(pin int, pinValue int) (err error)

	GetPinValue(pin int) (pinValue int, err error)

	CleanUp() (err error)
}

type gpio struct {
	pin_exporter   *os.File
	pin_unexporter *os.File

	initialized bool

	initializedPins map[int]pin
}

type pin struct {
	currentDir string
	currentVal int

	dir        *os.File
	val        *os.File
	active_low *os.File
	edge       *os.File
}

func (io *gpio) init() (err error) {
	if io.initialized {
		return
	}

	if io.pin_exporter, err = os.OpenFile("/sys/class/gpio/export", os.O_WRONLY, os.ModeExclusive); err != nil {
		return errors.New(fmt.Sprintf("Gpio: cannot open export file due to %s", err.Error()))
	}
	if io.pin_unexporter, err = os.OpenFile("/sys/class/gpio/unexport", os.O_WRONLY, os.ModeExclusive); err != nil {
		return errors.New(fmt.Sprintf("Gpio: cannot open unexport file due to %s", err.Error()))
	}

	io.initialized = true
	return
}

func (io *gpio) SetPinDir(pin int, pinDir string) (err error) {
	if _, allowed := allowedPins[pin]; allowed {
		if allowedPins[pin] == "used" {
			log.Panicf("Gpio: pin %i is used by board peripherals. Using it as gpio may cause other functionalitites to stop", pin)
		}
		io.init()
		currentPin := io.initializedPins[pin]
		if currentPin.dir == nil {
			if _, err = io.pin_exporter.WriteString(strconv.Itoa(pin)); err != nil {
				return
			}

			pin_dir_file_path := fmt.Sprintf("/sys/class/gpio/gpio%i/direction", pin)
			if currentPin.dir, err = os.OpenFile(pin_dir_file_path, os.O_RDWR, os.ModeExclusive); err != nil {
				return
			}
		}

		if _, err = currentPin.dir.WriteString(pinDir); err != nil {
			return
		}

		currentPin.currentDir = pinDir
	} else {
		return errors.New(fmt.Sprintf("Gpio: pin %i is used or allocated and not available for io", pin))
	}
	return
}

func (io *gpio) SetPinValue(pin int, pinValue int) (err error) {
	initialized_pin := io.initializedPins[pin]
	if initialized_pin.dir == nil {
		return errors.New(fmt.Sprintf("Gpio: direction not set for pin %i", pin))
	}

	if initialized_pin.currentDir == In {
		return errors.New(fmt.Sprintf("Gpio: pin %i is set as INPUT. Cant write to it", pin))
	}

	if initialized_pin.val == nil {
		pin_val_file_path := fmt.Sprintf("/sys/class/gpio/gpio%i/value", pin)

		if initialized_pin.val, err = os.OpenFile(pin_val_file_path, os.O_RDWR, os.ModeExclusive); err != nil {
			return
		}
	}

	if _, err = initialized_pin.val.WriteString(strconv.Itoa(pinValue)); err != nil {
		return
	}
	initialized_pin.currentVal = pinValue
	return
}

func (io *gpio) GetPinValue(pin int) (pinValue byte, err error) {
	initialized_pin := io.initializedPins[pin]
	if initialized_pin.dir == nil {
		return Low, errors.New(fmt.Sprintf("Gpio: direction not set for pin %i", pin))
	}

	if initialized_pin.val == nil {
		pin_val_file_path := fmt.Sprintf("/sys/class/gpio/gpio%i/value", pin)

		initialized_pin.val, err = os.OpenFile(pin_val_file_path, os.O_RDWR, os.ModeExclusive)
		if err != nil {
			return
		}
	}
	var num int
	var byteArray = make([]byte, 10)
	if num, err = initialized_pin.val.Read(byteArray); err != nil {
		return
	}
	status := string(byteArray[:num])
	if status == strconv.Itoa(1) {
		return High, nil
	} else {
		return Low, nil
	}
}

func (io *gpio) CleanUp() {
	for pinNum, _ := range io.initializedPins {
		io.pin_unexporter.WriteString(strconv.Itoa(pinNum))
	}
}
