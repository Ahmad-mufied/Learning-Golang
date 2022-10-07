package main

import "fmt"
/*
Variable

-- Untuk tempat untuk menyimpan data
-- Digunakan agar kita bisa mengakses data yang sama dimanapun kita mau
-- Di Go-Lang Variable hanya bisa menyimpan tipe data yang sama, jika kita ingin menyimpan data yang berbeda-beda jenis,
   kita harus membuat beberapa variable
-- Untuk membuat variable, kita bisa menggunakan kata kunci var, lalu diikuti dengan nama variable dan tipe datanya
*/


func main() {
	var name string

	name = "Ahmad Mufied"
	println(name)

	name = "Ahmad N"
	println(name)

	/* Tipe Data Variable

	-- Saat kita membuat variable, maka kita wajib menyebutkan tipe data variable tersebut
	-- Namun jika kita langsung menginisialisasikan data pada variable nya, maka kita tidak wajib
	   menyebutkan tipe data variable nya
	*/

	var friendName = "Budi"
	fmt.Println(friendName)

	var age = 30
	fmt.Println(age)

	/* Kata Kunci Var
	
	-- Di Go-Lang, kata kunci var saat membuat variable nggak wajib.
	-- Yang penting saat membuat variable kita langsung menginisialisasi datanya.
	-- Supaya nggak menggunakan kata kunci var, kita harus menggunakan kata kunci := saat
		menginisialisasikan data pada variable tersebut
	*/

	country := "Indonesia"
	fmt.Println(country)

	/* Deklarasi Multiple Variable
	-- Di Go-Lang kita bisa membuat variable sekaligus banyak
	-- Code yang dibuat akan lebih bagus dan mudah dibaca

	*/

	var (
		firstName = "Ahmad"
		lastName = "Mufied"
	)

	fmt.Println(firstName)
	fmt.Println(lastName)
}