package main

/*
Package & Import

--	Package

	--	Package adalah tempat yang bisa digunakan untuk mengorganisir kode program yang kita buat di
		Go-Lang
	--	Dengan menggunakan package, kita bisa merapikan kode program yang kita buat
	--	Package sendiri sebenarnya hanya direktori folder di sistem operasi kita

--	Import
	
	--	Secara standar, file Go-Lang hanya bisa mengakses file Go-Lang lainnya yang berada dalam
		package yang sama
	--	Jika kita ingin mengakses file Go-Lang yang berada diluar package, maka kita bisa menggunakan
		Import
*/

import (
	"1-Golang-Dasar/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
	fmt.Println("")
}