package main

import "fmt"

/*
* If Expression

-- If adalah salah satu kata kunci yang digunakan untuk percabangan
-- Percabangan artinya kita bisa mengeksekusi kode program tertentu ketika suatu kondisi terpenuhi
-- Hampir di semua bahasa pemrograman mendukung if expression


 * Else Expression

-- Blok if akan dieksekusi ketika kondisi if bernilai true
-- Kadang kita ingin melakukan eksekusi program tertentu jika kondisi if bernilai false
-- Hal ini bisa dilakukan menggunakan else expression

* Else If Expression

-- Kadang dalam If, kita butuh membuat beberapa kondisi
-- Kasus seperti ini, kita bisa menggunakan Else If expression

*/

func main() {
	var name = "Nugroho"

	if name == "Ahmad" {
		fmt.Println("Hello Ahmad")
	} else if name == "Nugroho"{
		fmt.Println("Hello Nugroho")
	} else {
		fmt.Println("HI,kenalan donk")
	}

	/*
	* If dengan Short Statement
	
	-- If mendukung short statement sebelum kondisi
	-- Hal ini sangat cocok untuk membuat statement yang sederhana sebelum melakukan pengecekan
	   terhadap kondisi
	*/

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama sudah benar")
	}

}