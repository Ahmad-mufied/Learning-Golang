package main

import "fmt"

/*
Type Assertions

--	Type Assertions merupakan kemampuan merubah tipe data menjadi tipe data yang diinginkan
--	Fitur ini sering sekali digunakan ketika kita bertemu dengan data interface kosong
*/

func random() interface{} {
	return "Ups"
} 

func main() {
	var result interface{} = random()
	var resultString string = result.(string)
	fmt.Println(resultString)

	/*
	Type Assertions Menggunakan Switch

	--	Saat salah menggunakan type assertions, maka bisa berakibat terjadi panic di aplikasi kita
	--	Jika panic dan tidak ter-recover, maka otomatis program kita akan mati
	--	Agar lebih aman, sebaiknya kita menggunakan switch expression untuk melakukan type assertions
	*/
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "is string")
	case int:
		fmt.Println("Value", value, "is int")
	default:
		fmt.Println("Unknown type")
	}
}