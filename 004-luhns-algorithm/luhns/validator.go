package luhns

import (
	"fmt"
	"strconv"
	"strings"
)

func Validator(input string) (valid bool, err error) {
	/* Step 1: 	Starting from the rightmost digit,
	*			double the value of every second digit
	*			If the double value is 2 digits (>9),
	* 			add its digits to get the final value
	*			This doubles should replace the values in
	* 			the positions in the initial array
	*
	* Step 2:	Get the sum of all values in the array
	*
	* Step 3:	If the sum % 10 is 0, then the number is valid
	 */

	var intVal int
	var inputArr = strings.Split(input, "")
	for i := len(inputArr) - 2; i >= 0; i -= 2 {
		intVal, err = strconv.Atoi(inputArr[i])
		if err != nil {
			fmt.Printf("input contains an invalid format")
			return
		}
		double := doubleNumber(intVal)
		inputArr[i] = fmt.Sprintf("%d", double)
	}

	sum := getSum(inputArr)
	return sum%10 == 0, nil
}

func doubleNumber(val int) (double int) {
	double = val * 2
	if double > 9 {
		ones := double - 10
		return ones + 1
	}
	return double
}

func getSum(inputArr []string) (sum int) {
	for _, val := range inputArr {
		intVal, _ := strconv.Atoi(val)
		sum += intVal
	}
	return sum
}
