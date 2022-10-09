package main

import "fmt"

/*
	Tipe Data Slice

-- Tipe data Slice adalah potongan dari data Array
-- Slice mirip dengan Array, yang membedakan adalah ukuran Slice bisa berubah
-- Slice dan Array selalu terkoneksi, dimana Slice adalah data yang mengakses sebagian
   atau seluruh data di Array


   Detail Tipe Data Slice

-- Tipe Data Slice memiliki 3 data, yaitu pointer, length, dan capacity
-- Pointer adalah penunjuk data pertama di array pada slice
-- Length adalah panjang dari slice
-- Capacity adalah kapasitas dari slice, dimana length tidak boleh lebih dari capacity
*/

func main() {
	var months = [...]string{
		"January",
		"February",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	/*
	Membuat Slice dari Array
	
		Membuat Slice					Keterangan
	array[low:high]			Membuat slice dari array dimulai index low sampai index sebelum high
	array[low:]				Membuat slide dari array dimulai index low sampai index akhir di array
	array[:high]			Membuat slice dari array dimulai index 0 sampai index sebelum high
	array[:]				Membuat slice dari array dimulai index 0 sampai index akhir di array
	*/

	var slice1 = months[4:7]
	fmt.Println(slice1)
	fmt.Println(len(slice1))
	fmt.Println(cap(slice1))

	months[5] = "Diubah"
	fmt.Println(slice1)

	slice1[0] = "Mei Lagi"
	fmt.Println(months)

	var slice2 = months[8:11]
	fmt.Println(slice2)

	/*
	Function Slice
	
		Operasi								Keterangan
	len(slice)								Mendapatkan panjang.
	cap(slice)								Mendapatkan kapasitas.
	append(slice, data)						Membuat slice baru dengan menambah data ke
											posisi terakhir slice, jika kapasitas sudah
											penuh, maka akan membuat array baru.
	make([]TypeData, length, capacity)		Membuat slice baru
	copy(destination, source)				Menyalin slice dari source ke destination
	*/
	

	// Kode Program Append Slice

	var slice3 = append(slice2, "Eko")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// Kode Program Make Slice

	newSlice := make([]string, 3, 5)
	newSlice[0] = "Ahmad"
	newSlice[1] = "Mufied"

	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// Kode Program Copy Slice

	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	// Kode Program Array vs Slice

	iniArray := [5]int{1, 2, 3, 4, 5}
	iniSlice := []int{1, 2, 3, 4, 5}

	fmt.Println(iniArray)
	fmt.Println(iniSlice)
	
}