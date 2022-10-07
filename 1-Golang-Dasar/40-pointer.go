package main

import "fmt"

type Address struct {
	City, Province, Country string
}

func main() {

	fmt.Println("==== Pass by Value =====")
	alamat1 := Address{"Subang", "Jawa Barat", "Indonesia"}
	alamat2 := alamat1

	alamat2.City = "Bandung"

	fmt.Println(alamat1)
	fmt.Println(alamat2)
	fmt.Println()

	fmt.Println("==== Pass by Reference ====")

	var address1 Address = Address{"Subang", "Jawa Barat", "Indonesia"}
	// var address4 *Address = &Address{"Subang", "Jawa Barat", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

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

	fmt.Println("===== Function new ======")
	//* Untuk membuat pointer dan mengembalikan pointer ke data kosong, artinya tidak ada data awal
	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)
	

	

	
}