package gpio

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type analogPin struct {
	n int

	val         *os.File
	initialized bool
}

func newAnalogPin(n int) (*analogPin, error) {
	p = &analogPin{n: n}
	if err := p.init(); err != nil {
		return nil, err
	}
	return p, nil
}

func (p *analogPin) init() error {
	if f, err := os.OpenFile("/sys/devices/bone_capemgr.8/slots", os.O_RDWR, os.ModeExclusive); err != nil {
		return
	}

	if !initialized(f) {
		if err := f.WriteString("cape-bone-iio"); err != nil {
			return
		}
	}

	fileName := fmt.Sprintf("/sys/devices/ocp.2/helper.14/AIN%v", p.n)
	if p.val, err = os.OpenFile(fileName, os.O_RDONLY, os.ModeExclusive); err != nil {
		return
	}

	return nil
}

func initialized(f *os.File) bool {
	buf := make([]byte, 300)
	if count, err := f.Read(buf); err != nil {
		panic(err)
	}
	f.val.Seek(0, 0)

	s := string(buf[:count])

	return strings.Contains(s, "cape-bone-iio")
}

func (p *analogPin) Read() (int, error) {
	buf := make([]byte, 5)
	if count, err := p.val.Read(buf); err != nil {
		return nil, err
	}
	p.val.Seek(0, 0)

	s := string(buf[:count])
	s = strings.TrimSpace(s)

	return strconv.Atoi(s), nil
}

func (p *analogPin) Close() error {
	if err := p.val.Close(); err != nil {
		return err
	}
	return nil
}
