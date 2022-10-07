package main

import (
	"fmt"
	"strconv"
)

func main() {
	boolean, err := strconv.ParseBool("true")

	if err == nil {
		fmt.Println(boolean)
	} else {
		fmt.Println(err.Error())
	}

	number, err := strconv.ParseInt("1000000", 10, 64)
	if err == nil {
		fmt.Println(number)
	} else {
		fmt.Println(err.Error())
	}

	
	//* Atoi is equivalent to ParseInt(s, 10, 0), converted to type int.
	number1, err := strconv.Atoi("1233")
	if err == nil {
		fmt.Println(number1)
	} else {
		fmt.Println(err.Error())
	}

	//* Itoa is equivalent to FormatInt(int64(i), 10). 
	number2 := strconv.Itoa(12)
	fmt.Println(number2)

	//* int to string
	value := strconv.FormatInt(1000000, 10)
	fmt.Println(value)

	//* int to binnary
	binary := strconv.FormatInt(1000000, 2)
	fmt.Println(binary)

	//* int to octal
	binary1 := strconv.FormatInt(1000000, 2)
	fmt.Println(binary1)

	
}