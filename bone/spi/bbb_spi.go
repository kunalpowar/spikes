package main

import (
	"syscall"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Fetching slots file")
	slotsFilePath := "/sys/devices/bone_capemgr.8/slots"

	var err error
	var slotsFile *os.File
	var spiDev0File *os.File

	if slotsFile, err = os.OpenFile(slotsFilePath, os.O_RDWR, os.ModeExclusive); err != nil {
		panic(err)
	}

	fmt.Println("Fetched slots file")

	slotsFile.WriteString("s")

	fmt.Println("Reading spidev file")

	spiDev0Path := "/dev/spidev1.0"

	if spiDev0File, err = os.OpenFile(spiDev0Path, os.O_RDWR, os.ModeExclusive); err != nil {
		panic(err)
	}
	defer slotsFile.Close()
	defer spiDev0File.Close()

	fmt.Println("Read spi dev file")

	spiDev0FileFd := spiDev0File.Fd()

	fmt.Printf("the spidev's file fd is: %v", spiDev0FileFd)

	syscall.Syscall(syscall.SYS_IOCTL, a1, a2, a3)
}
