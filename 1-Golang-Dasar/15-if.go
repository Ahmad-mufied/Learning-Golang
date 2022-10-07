package main

import "fmt"

func main() {
	var name = "Nugroho"

	if name == "Ahmad" {
		fmt.Println("Hello Ahmad")
	} else if name == "Nugroho"{
		fmt.Println("Hello Nugroho")
	} else {
		fmt.Println("HI,kenalan donk")
	}

	// IF Short Statement

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}