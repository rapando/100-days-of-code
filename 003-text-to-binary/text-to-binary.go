package main

import (
	"fmt"
	"os"
)

func main() {
	var binaryOutput string
	fmt.Printf("This program converts text to binary.\n")
	args := os.Args[1:]
	if len(args) != 1 {
		fmt.Println(`
		Only one argument is required.
		Try with a string or read the documentation at : 
		https://github.com/rapando/100-days-of-code/tree/main/003-text-to-binary
		
		Cheers!`)
		return
	}

	text := args[0]
	textBytes := []byte(text)
	for _, b := range textBytes {
		binaryOutput += intToBinary(int(b))
	}

	fmt.Printf("\nInput   %s\n", text)
	fmt.Printf("\nOutput \n%s\n", binaryOutput)
}

func intToBinary(n int) (binaryString string) {
	var binary []int
	if n == 0 {
		return eightBitBinary("0")
	}

	// while n/2 is not 0, append the value of n%2 to binary array
	for n != 0 {
		binary = append(binary, n%2)
		n /= 2
	}

	// print the binary array backwards
	for i := len(binary) - 1; i >= 0; i-- {
		binaryString += fmt.Sprintf("%d", binary[i])
	}

	return eightBitBinary(binaryString)
}

func eightBitBinary(binary string) string {
	// if the binary representation has 7 or less characters
	var prefix = ""
	length := len(binary)
	deficit := 8 - length
	for i := 0; i < deficit; i++ {
		prefix += "0"
	}
	return prefix + binary + " "
}
