package main

import "fmt"

/*
Operasi Boolean

Operator			Keterangan
	&&					Dan
	||					Atau
	!					Kebalikan
*/

func main() {
	var ujian = 80
	var absensi = 80

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)

	var lulus = lulusAbsensi && lulusUjian
	fmt.Println(lulus)

	fmt.Println(ujian >= 80 && absensi >= 80)
}