package main

import "fmt"

func main() {
	name := "Ahmad"
	counter := 0

	increament := func() {
		name := "Mufied"
		fmt.Println("Increament")
		counter ++
		fmt.Println(name)
	}

	increament()
	increament()

	fmt.Println(counter)
	fmt.Println(name)
}