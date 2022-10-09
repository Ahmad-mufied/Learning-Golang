package main

import "fmt"

/*
Returning Multiple Values

--	Function tidak hanya dapat mengembalikan satu value, tapi juga bisa multiple value
--	Untuk memberitahu jika function mengembalikan multiple value, kita harus menulis semua tipe
	data return value nya di function
*/

func getFullName() (string, string, string) {
	return "Ahmad", "Mufied", "Nugroho"
}

func main() {

	/*
	Menghiraukan Return Value

	--	Multiple return value wajib ditangkap semua value nya
	--	Jika kita ingin menghiraukan return value tersebut, kita bisa menggunakan tanda _ (garis bawah)
	*/
	fisrtName, _, lastName := getFullName()
	fmt.Println(fisrtName)
	fmt.Println(lastName)
}