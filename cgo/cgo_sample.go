package main

/*
#include <stdlib.h>
#include <sys/ioctl.h>
*/
import "C"
import "fmt"

func Random() int {
	return int(C.random())
}

func Seed(i int) {
	C.srandom(C.uint(i))
}

func main() {
	fmt.Println(Random())
	fmt.Println(C.TIOCGSIZE)
}
