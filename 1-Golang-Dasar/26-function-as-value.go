package main

import "fmt"

/*
Function Value

--	Function adalah first class citizen
--	Function juga merupakan tipe data, dan bisa disimpan di dalam variable
*/

func getGoodBye(name string) string {
	return "Good Bye " + name
}

func main() {
	sayGoodBye := getGoodBye

	result := sayGoodBye("Ahmad")
	fmt.Println(result)
}