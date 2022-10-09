package main

import "fmt"

/*
Named Return Values

-- Biasanya saat kita memberi tahu bahwa sebuah function mengembalikan value, maka kita hanya
	mendeklarasikan tipe data return value di function
--	Namun kita juga bisa membuat variable secara langsung di tipe data return function nya
*/

func getFullName2() (firstName string, middleName string, lastName string) {
	firstName = "Ahmad"
	middleName = "Mufied"
	lastName = "Nugroho"
	return
}

func main() {
	a, b, c := getFullName2()
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
}