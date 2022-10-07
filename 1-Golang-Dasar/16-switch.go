package main

import "fmt"

func main() {
	name := "Mufied Nugroho"

	switch name {
	case "Ahmad":
		fmt.Println("Hello Mufied")
		fmt.Println("Hello Mufied")
	case "Mufied":
		fmt.Println("Hello Mufied")
		fmt.Println("Hello Mufied")
	default:
		fmt.Println("Kenalan Donk")
		fmt.Println("Kenalan Donk")
	}

	// Short Statement
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang.")
	case false:
		fmt.Println("Nama sudah benar.")
	}

	// Switch tanpa kondisi
	panjangNama := len(name)
	switch {
	case panjangNama > 10:
		fmt.Println("Nama terlalu panjang")
	case panjangNama > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

}