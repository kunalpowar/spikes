package main

import (
	"fmt"
	"os"
	// "strconv_"
	"syscall"
	"time"
	"unsafe"
)

func main() {
	fmt.Println("Reading file")
	file, _ := os.Open("/dev/cu.usbmodemfd121")
	// file, _ := os.OpenFile("/dev/cu.usbmodemfd121", os.O_EXCL, os.ModeExclusive)
	fd := file.Fd()
	// fd, _ := syscall.Open("/dev/tty.usbmodemfd121", syscall.O_EXCL|syscall.O_NONBLOCK, syscall.SYS_READ)
	time.Sleep(5 * time.Second)

	var serial_data termios
	read_data := make([]byte, 100)
	data := make([]byte, 1)
	data[0] = 1

	fmt.Printf("\nFile FD is %v", fd)

	_, _, errno := syscall.Syscall(syscall.SYS_IOCTL, fd, syscall.TIOCGETA, uintptr(unsafe.Pointer(&serial_data)))

	fmt.Printf("\nRead using syscall with errorno %v", errno)
	fmt.Printf("\nspeed: %v", &serial_data)

	fmt.Printf("\nThe fd is %v", fd)

	// num, err := file.Write(data)
	// num, err := syscall.Write(fd, data)
	// fmt.Printf("\n %v bytes written to the file", num)
	// fmt.Println(err.Error())

	// syscall.Write(fd, data)

	// num, err := file.Read(data)
	// if err != nil {
	// 	fmt.Printf("Error in file write", err.Error())
	// }

	for i := 0; i < 100; i++ {
		_, err := file.Read(read_data)
		if err != nil {
			fmt.Printf("Error in file Read", err.Error())
		}
		// num, err = syscall.Read(fd, read_data)
		fmt.Printf("\nRead %v", read_data[0])
	}

}

// syscall.Syscall(syscall.SYS_IOCTL, b.fd, SPI_IOC_RD_MAX_SPEED_HZ, uintptr(unsafe.Pointer(&speed_max)))

type termios struct {
	c_iflag  uint64
	c_oflag  uint64
	c_cflag  uint64
	c_lflag  uint64
	c_cc     [20]byte
	c_ispeed uint64
	c_ospeed uint64
}
