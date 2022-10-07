package main

import "fmt"

/* Constant

-- Constant adalah variable yang nilainya tidak bisa diubah lagi setelah pertama kali diberi nilai
-- Cara pembuatan constant mirip dengan variable, yang membedakan hanya kata kunci yang digunakan adalah const, bukan var
-- Saat pembuatan constant, kita wajib langsung menginisialisasikan datanya.
*/

func main() {
	const firstName string = "Ahmad"
	const lastName = "Mufied"
	const value = 1000

	/* Multitple const
	-- Sama seperti variable, di Go-Lang  juga kita bisa membuat constant secara sekaligus banyak
	 */
	const (
		namaDepan string = "Ahmad"
		namaBelakang = "Mufied"
		nilai = 1000
	)

	fmt.Println(namaDepan)
	fmt.Println(namaBelakang)
}