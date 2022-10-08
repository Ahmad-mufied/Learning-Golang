package main

import "fmt"

/*
Konversi Tipe Data

-- Di Go-Lang kadang kita butuh melakukan konversi tipe data dari satu tipe ke tipe lain
-- Misal kita ingin mengkonversi tipe data int32 ke int63, dan lain-lain

*/

func main() {
	var nilai32 int32 = 100000
	var nilai64 int64 = int64(nilai32)
	var nilai8 int8 = int8(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai8)

	var name = "Eko"
	var e byte = name[0]
	var eString string = string(e)

	fmt.Println(name)
	fmt.Println(eString)
}