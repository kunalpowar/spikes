package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	setUpAnalogPins()
	// findAnalogInputFile()
}

var file *os.File

func setUpAnalogPins() {
	f, err := os.OpenFile("/sys/devices/bone_capemgr.8/slots", os.O_RDWR, os.ModeExclusive)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(initialized(f))
	// _, err = f.WriteString("cape-bone-iio")
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// // fileName := fmt.Sprintf("/sys/devices/ocp.2/helper.14/AIN1", p.n)

	// file, err = os.OpenFile("/sys/devices/ocp.2/helper.14/AIN1", os.O_RDONLY, os.ModeExclusive)
	// if err != nil {
	// 	panic(err)
	// }

}
func initialized(f *os.File) bool {
	buf := make([]byte, 300)
	count, err := f.Read(buf)
	if err != nil {
		panic(err)
	}
	f.Seek(0, 0)

	s := string(buf[:count])

	return strings.Contains(s, "cape-bone-iio")
}

func findAnalogInputFile() {
	var err error
	var count int
	for i := 0; i < 20; i++ {
		time.Sleep(1 * time.Second)

		buf := make([]byte, 5)
		count, err = file.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
		}

		file.Seek(0, 0)

		s := string(buf[:count])
		s = strings.TrimSpace(s)
		i, _ := strconv.Atoi(s)
		fmt.Printf("lighting: %v\n", i)
	}
}
