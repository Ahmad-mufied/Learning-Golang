package main

import "fmt"

/*
* For Loops
-- Dalam bahasa pemrograman, biasanya ada fitur yang bernama perulangan
-- Salah satu fitur perulangan adalah for loops
*/

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	/*
	* For dengan Statement
	-- Dalam for, kita bisa menambahkan statement, dimana terdapat 2 statement yang bisa tambahkan
	   di for
	-- Init statement, yaitu statement sebelum for di eksekus
	-- Post statement, yaitu statement yang akan selalu dieksekusi di akhir tiap perulangan
	 */
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	fmt.Println()


	slice := []string{"Ahmad", "Mufied", "Nugroho"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	/*
	* For Range
	-- For bisa digunakan untuk melakukan iterasi terhadap semua data collection
	-- Data collection contohnya Array, Slice dan Map
	*/

	// * For range slice

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	// * For range Map

	person := make(map[string]string)
	person["name"] = "Ahmad"
	person["tittle"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}