package main

import "fmt"

/*
Closures

--	Closure adalah kemampuan sebuah function berinteraksi dengan data-data disekitarnya dalam scope yang sama
--	Harap gunakan fitur closure ini dengan bijak saat kita membuat aplikasi
*/

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