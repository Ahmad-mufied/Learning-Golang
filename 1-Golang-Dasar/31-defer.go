package main

import "fmt"

/*

Defer

--	Defer function adalah function yang bisa kita jadwalkan untuk dieksekusi setelah sebuah function
	selesai di eksekusi
--	Defer function akan selalu dieksekusi walaupun terjadi error di function yang dieksekusi
*/

func logging() {
	fmt.Println("Selesai memanggil function")
}

func runApplication(value int) {
	defer logging()
	fmt.Println("Run application")
	result := 10 / value
	fmt.Println("Result ", result)
}

func main()  {
	runApplication(0)
}

