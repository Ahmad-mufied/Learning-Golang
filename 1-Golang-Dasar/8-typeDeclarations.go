package main

import "fmt"

/*
Type Declarations
-- Kemampuan membuat ulang tipe data baru dari tipe data yang sudah ada
-- iasanya digunakan untuk membuat alias terhadap tipe data yang sudah ada,
dengan tujuan agar lebih mudah dimengerti
*/

func main() {
	type NoKTP string
	type Married bool

	var noKtpEko NoKTP = "234324234234324324324324"
	var marriedStatus Married = true
	fmt.Println(noKtpEko)
	fmt.Println(marriedStatus)
}