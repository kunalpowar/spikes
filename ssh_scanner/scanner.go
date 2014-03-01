package main

import (
	"fmt"
	"os/exec"
)

func execOutput(name string, arg ...string) (output string, err error) {
	var out []byte
	if out, err = exec.Command(name, arg...).Output(); err != nil {
		return
	}
	output = string(out)
	return
}

func main() {
	out, err := execOutput("ssh","root@192.168.2.12")

	fmt.Print(out)
	fmt.Print(err.Error())
}
