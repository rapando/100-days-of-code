package main

import (
	"fmt"

	"github.com/rapando/100-days-of-code/004-luhns-algorithm/luhns"
)

func main() {
	input := "7992739813"
	valid, err := luhns.Validator(input)
	if err != nil {
		fmt.Printf("%s has invalid characters. Only digits are allowed\n", input)
		return
	}

	if valid {
		fmt.Printf("%s is valid\n", input)
	} else {
		fmt.Printf("%s is NOT valid\n", input)
	}
}
