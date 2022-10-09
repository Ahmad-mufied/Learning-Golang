package main

 import "fmt"

 /*
 Pointer di Function

 --	Saat kita membuat parameter di function, secara default adalah pass by value, artinya data akan di
	copy lalu dikirim ke function tersebut
--	Oleh karena itu, jika kita mengubah data di dalam function, data yang aslinya tidak akan pernah
	berubah.
--	Hal ini membuat variable menjadi aman, karena tidak akan bisa diubah
--	Namun kadang kita ingin membuat function yang bisa mengubah data asli parameter tersebut
--	Untuk melakukan ini, kita juga bisa menggunakan pointer di function
--	Untuk menjadikan sebuah parameter sebagai pointer, kita bisa menggunakan operator * di
	parameternya
 */

 type Address struct {
	City, Province, Country string
 }

 func ChangeAddressToIndonesia(address *Address) {
	address.Country = "Indonesia"
 }

func main() {

	var address1 Address = Address{"Bandar Lampung", "Lampung", "Indonesia"}
	var address2 *Address = &address1
	var address3 *Address = &address1

	address2.City = "Bandung"

	*address2 = Address{"Malang", "Jawa Timur", "Indonesia"}

	fmt.Println(address1)
	fmt.Println(address2)
	fmt.Println(address3)

	var address4 *Address = new(Address)
	address4.City = "Jakarta"
	fmt.Println(address4)

	var alamat = Address{
		City: "Subang",
		Province: "Jawa Barat",
		Country: "",
	}
	
	var alamatPointer *Address = &alamat
	ChangeAddressToIndonesia(alamatPointer)
	fmt.Println(alamat)
} 