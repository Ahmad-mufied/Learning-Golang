package main

import "fmt"

/*
Function Parameter

--	Saat membuat function, kadang-kadang kita membutuhkan data dari luar, atau kita sebut
	parameter.
--	Kita bisa menambahkan parameter di function, bisa lebih dari satu
--	Parameter tidaklah wajib, jadi kita bisa membuat function tanpa parameter seperti sebelumnya
	yang sudah kita buat
--	Namun jika kita menambahkan parameter di function, maka ketika memanggil function tersebut,
	kita wajib memasukkan data ke parameternya
*/

func sayHelloTo(firstName string, lastName string) {
	fmt.Println("Hello", firstName, lastName)
}

func main() {
	firstName := "Ahmad"
	sayHelloTo(firstName, "Mufied")
}