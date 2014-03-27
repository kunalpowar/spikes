package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	a, _ := findFirstMatchingFile("/Users/kunalpowar/Projects/hardware/hardware_spikes/*/analog_in")
	fmt.Println(a)
}

func findFirstMatchingFile(glob string) (string, error) {
	matches, err := filepath.Glob(glob)
	if err != nil {
		return "", err
	}
	if len(matches) >= 1 {
		return matches[0], nil
	}
	return "", nil
}
