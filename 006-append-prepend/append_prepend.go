package main

import (
	"fmt"
)

func main() {

	arr := []int{1, 2, 3, 5}
	fmt.Println(arr)

	// use the append function to append to an array
	// if you have not defined the size of the array, the size is the current number of items by default
	// in this case, the size is 4. If we try to append something to it it will not work because the array is full
	// but the function append will actually create a new array and then replace the old one with the new array. This is a very important concept to know
	// therefore, always use the append function when you want to append

	arr = append(arr, 34)
	fmt.Println(arr)

	// to prepend we can cleverly use append.
	// say we want to prepend 100.
	// 1. We create an array with the 100 as the only element
	// 2. We break down arr into its individual elements by adding three dots to it
	// 3. Now append the arr to the array in step one

	prependItem := []int{100}
	arr = append(prependItem, arr...)
	fmt.Println(arr)

	// So far we have now seen how to append/ prepend using the append function.
	// using the prepend trick, we can append/ prepend multiple items

	appendItems := []int{1, 1, 1, 1}
	arr = append(arr, appendItems...)
	fmt.Println(arr)

	prependItems := []int{8, 8, 8, 8}
	arr = append(prependItems, arr...)
	fmt.Println(arr)

	fmt.Println("Cheers")

}
