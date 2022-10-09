package main

import "fmt"

/*
* Switch Expression

-- Selain if expression, untuk melakukan percabangan, kita juga bisa menggunakan Switch Expression
-- Switch expression sangat sederhana dibandingkan if
-- Biasanya switch expression digunakan untuk melakukan pengecekan ke kondisi dalam satu
   variable
*/

func main() {
	name := "Mufied Nugroho"

	switch name {
	case "Ahmad":
		fmt.Println("Hello Mufied")
		fmt.Println("Hello Mufied")
	case "Mufied":
		fmt.Println("Hello Mufied")
		fmt.Println("Hello Mufied")
	default:
		fmt.Println("Kenalan Donk")
		fmt.Println("Kenalan Donk")
	}

	/*
	* Short Statement
	-- Sama dengan If, Switch juga mendukung short statement sebelum variable yang akan di cek
	   kondisinya
	 */
	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Nama terlalu panjang.")
	case false:
		fmt.Println("Nama sudah benar.")
	}

	/* 
	* Switch tanpa kondisi
	-- Kondisi di switch expression tidak wajib
	-- Jika kita tidak menggunakan kondisi di switch expression, kita bisa menambahkan kondisi tersebut
	   di setiap case nya
	   
	*/
	panjangNama := len(name)
	switch {
	case panjangNama > 10:
		fmt.Println("Nama terlalu panjang")
	case panjangNama > 5:
		fmt.Println("Nama lumayan panjang")
	default:
		fmt.Println("Nama sudah benar")
	}

}