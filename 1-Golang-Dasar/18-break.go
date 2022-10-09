package main

import "fmt"

/*
* Break & Continue
-- Break & continue adalah kata kunci yang bisa digunakan dalam perulangan
-- Break digunakan untuk menghentikan seluruh perulangan
-- Continue adalah digunakan untuk menghentikan perulangan yang berjalan, dan langsung 
   melanjutkan ke perulangan selanjutnya.
*/

func main() {
	for i := 0; i < 10; i++ {

		fmt.Println("Perulangan ke", i)
		if i == 5 {
			break
		}
	}
}