//* Defer
//* -- Function yang bisa kita jadwalkan untuk diekesekusi setelah sebuah function selesai dieksekusi
//* -- Akan selalu dieksekusi walaupun terjadi error di function yang diekesekusi

package main

import "fmt"

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

