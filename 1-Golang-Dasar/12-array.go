package main

import "fmt"
/*
Tipe Data  Array
-- tipe data yang berisikan kumpulan data dengan tipe yang sama
-- Saat membuat array, kita perlu menentukan jumlah data yang bisa ditampung oleh Array tersebut
-- Daya tampung Array tidak bisa bertambah setelah Array dibuat
*/

func main() {

	var names [3]string

	names[0] = "Ahmad"
	names[1] = "Mufied"
	names[2] = "Nugroho"
	
	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])
	
	/*
	Membuat Array Langsung
	-- Di Go-Lang kita juga bisa membuat Array secara langsung saat deklarasi variable
	*/

	var values = [3]int{
		90,
		95,
		80,
	}

	fmt.Println(values)
	fmt.Println(values[0])
	fmt.Println(values[1])
	fmt.Println(values[2])

	/*
	Function Array

			Operasi								Keterangan
	len(array)						Untuk mendapatkan panjang Array
	array[index]					Mendapat data di posisi
	array[index] = value			Mengubah data di posisi index

	*/

	fmt.Println(len(names))
	fmt.Println(len(values))
	
	var lagi [10]string
	fmt.Println(len(lagi))

	var angka [5]int
	fmt.Println()
	angka[0] = 1
	angka[1] = 2
	fmt.Println(angka)

	angka1 := [10]int{
		1,
		2,
		3,
		4,
		5,
	}
	fmt.Println("len angka1 = ", len(angka1))
	fmt.Println("array angka1 = ", angka1)
}