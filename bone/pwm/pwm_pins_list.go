package main

import (
	"log"
	"os"
	// "time"
	"os/exec"
	"fmt"
)

func main() {
	fmt.Println("Started testing PWM on beagle bone")

	runCmd("modprobe", "pwm_test")

	runCmd("find", "/ -name *bone_pwm_P*.dts")
}

func runCmd(cmd, options string) {
	fmt.Printf("pwmTest: running command [%s %s]", cmd, options)
	out, err := exec.Command(cmd,options).Output()
	if err != nil {
		log.Fatalf("Error: %s", err.Error())
	}
	os.Stdout.Write(out)
	// fmt.Printf("pwmTest: running command [%s %s]", cmd, options)
	// c := exec.Command(cmd,options)
	
	// var out []byte

	// out, _ = c.Output()

	// time.Sleep(5 * time.Second)
	// // if err != nil {
	// // 	fmt.Errorf("pmwTest: command %s falied due to %v\n", cmd ,err.Error())
	// // }
	
	// fmt.Println(string(out))	
	// fmt.Printf("pwmTest: command [%s %s] run without errors\n", cmd, options)
}
