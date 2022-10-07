package main

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpEko NoKTP = "234324234234324324324324"
	var marriedStatus Married = true
	fmt.Println(noKtpEko)
	fmt.Println(marriedStatus)
}