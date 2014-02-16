package main

import (
	"fmt"
)

func main() {
	a := map[uint]string{
		1: "haha",
		2: "bleh",
	}

	_, val := a[1]
	fmt.Println(val)

	_, val = a[2]
	fmt.Println(val)

	_, val = a[3]
	fmt.Println(val)

}
