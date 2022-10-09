package main

import "fmt"

/*
Pointer
*/

type Address struct {
	City, Province, Country string
}

func main() {

	/*
	Pass by Value

	--	Secara default di Go-Lang semua variable itu di passing by value, bukan by reference
	--	Artinya, jika kita mengirim sebuah variable ke dalam function, method atau variable lain,
		sebenarnya yang dikirim adalah duplikasi value nya
	*/
	fmt.Println("==== Pass by Value =====")
	alamat1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	alamat2 := alamat1

	alamat2.City = "Bandung"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
	fmt.Println()

	/*
	Pass by Reference dengan Pointer

	--	Pointer adalah kemampuan membuat reference ke lokasi data di memory yang sama, tanpa menduplikasi 
		data yang sudah ada
	--	Sederhananya, dengan kemampuan pointer, kita bisa membuat pass by reference

	
	Operator &
	--	Untuk membuat sebuah variable dengan nilai pointer ke variable yang lain, kita bisa menggunakan
		operator & diikuti dengan nama variable nya
	*/

	fmt.Println("==== Pass by Reference ====")

	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	// var address4 *Address = &Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	/*
	Operator *

	--	Saat kita mengubah variable pointer, maka yang berubah hanya variable tersebut.
	--	Semua variable yang mengacu ke data yang sama tidak akan berubah
	--	Jika kita ingin mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan
		operator *
	*/

	//* Mengubah seluruh variable yang mengacu ke data tersebut, kita bisa menggunakan operator *
	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)
	fmt.Println()

	//* Mengubah variable pointer, maka yang berubah hanya variable tersebut.
	address2 = &Address{"Natar", "Lampung Selatan", "Indonesia"}
	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println()

	/*
	Function new
	--	Sebelumnya untuk membuat pointer dengan menggunakan operator &
	--	Go-Lang juga memiliki function new yang bisa digunakan untuk membuat pointer
	--	Namun function new hanya mengembalikan pointer ke data kosong, artinya tidak ada data awal
	*/
	fmt.Println("===== Function new ======")
	//* Untuk membuat pointer dan mengembalikan pointer ke data kosong, artinya tidak ada data awal
	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)
	

	

	
}