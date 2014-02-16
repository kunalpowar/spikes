package main

import (
	"fmt"
	"os"
	// "strconv"
	// "time"
)

func main() {
	fmt.Println("Initializing Blinker")

	pin_exporter, err := os.OpenFile("/sys/class/gpio/export", os.O_WRONLY, os.ModeExclusive)

	if err != nil {
		fmt.Println("Error opening file: export")
		fmt.Println(err.Error())
	}

	// var num int
	_, err = pin_exporter.WriteString("31")
	if err != nil {
		fmt.Println(err.Error())
	}

	pin_val, err := os.OpenFile("/sys/class/gpio/gpio31/value", os.O_RDWR, os.ModeExclusive)
	if err != nil {
		fmt.Println("Error opening file: gpio31")
		fmt.Println(err.Error())
	}

	var num int
	a := make([]byte, 10)

	if num, err = pin_val.Read(a); err != nil {
		fmt.Printf(err.Error())
	}

	fmt.Println(string(a[:num]))

	// for i := 0; i < 10; i++ {
	// 	time.Sleep(1 * time.Second)
	// 	pin_val, err := os.OpenFile("/sys/class/gpio/gpio31/direction", os.O_RDWR, os.ModeExclusive)
	// 	if err != nil {
	// 		fmt.Println("Error opening file: gpio31")
	// 		fmt.Println(err.Error())
	// 	}
	// 	pin_val, err := os.OpenFile("/sys/class/gpio/gpio31/value", os.O_RDWR, os.ModeExclusive)
	// 	if err != nil {
	// 		fmt.Println("Error opening file: gpio31")
	// 		fmt.Println(err.Error())
	// 	}

	// 	var num int
	// 	num, err = pin_val.WriteString("out")
	// 	if err != nil {
	// 		fmt.Println(num)
	// 		fmt.Println(err.Error())
	// 	}

	// 	val := i % 2
	// 	num, err = pin_val.WriteString(strconv.Itoa(val))
	// 	if err != nil {
	// 		fmt.Println(num)
	// 		fmt.Println(err.Error())
	// 	}
	// }

}
