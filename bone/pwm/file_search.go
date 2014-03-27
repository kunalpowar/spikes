package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
	find()
}

func find() {
	cmd := exec.Command("find", "/Users/kunalpowar/Projects/hardware/hardware_spikes/bone", "-name", "analog*")
	out, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}
	fmt.Println(out)
	s := string(out)
	a := strings.Split(s, strings.Spli)
	fmt.Println(string(out))
}
