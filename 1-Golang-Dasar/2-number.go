package main

import "fmt"

/*
1. Tipe Data Number

-- 1. Integer

	int8		nilai min = -128					nilai maximum = 127
	int16 		nilai min = -32768					nilai maximum = 32767
	int32 		min = -2147483648					nilai maximum = 2147483647
	int64		min = -9223372036854775808			nilai maximum = 9223372036854775807

	uint8		nilai min = 0						nilai maximum = 255
	uint16		nilai min = 0						nilai maximum = 65535
	uint32		nilai min = 0						nilai maximum = 4294967295
	uint		nilai min = 0						nilai maximum = 18446744073709551615

-- 2. Floating Point

	float32 	nilai minimum = 1.18×10**-38		nilai maximum = 3.4×10**38
	float54 	nilai minimum = 2.23×10**−308		nilai maximum = 1.80×10**308

	complex64 	complex numbers with float32 real and imaginary parts.
	complex128	complex numbers with float64 real and imaginary parts.

----- Alias -----

byte = uint8
rune = int32
int = Minimal int32
uint = minimall int32
*/


func main() {
	fmt.Println("Satu = ", 1)
	fmt.Println("Dua = ", 2)
	fmt.Println("Tiga Koma Lima ", 3.5)
}