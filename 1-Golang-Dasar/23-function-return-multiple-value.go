package main

import "fmt"

func getFullName() (string, string, string) {
	return "Ahmad", "Mufied", "Nugroho"
}

func main() {
	fisrtName, _, lastName := getFullName()
	fmt.Println(fisrtName)
	fmt.Println(lastName)
}