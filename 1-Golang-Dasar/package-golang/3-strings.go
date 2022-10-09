package main

import (
	"fmt"
	"strings"
)

/*
Package strings

--	Package strings adalah package yang berisikan function-function untuk memanipulasi tipe data
	String
--	https://golang.org/pkg/strings/
*/



func main() {
	//*  strings.Contains
	fmt.Println(strings.Contains("Ahmad Mufied", "Ahmad"))
	fmt.Println(strings.Contains("Ahmad Mufied", "Budi"))

	//*  strings.Split
	fmt.Println(strings.Split("Ahmad Mufied Nugroho", " "))

	fmt.Println(strings.ToLower("Ahmad Mufied Nugroho"))
	fmt.Println(strings.ToUpper("Ahmad Mufied Nugroho"))
	fmt.Println(strings.Title("ahmad mufied nugroho"))

	//* strings.Trim
	fmt.Println(strings.Trim("    Ahmad Mufied      ", " "))

	//* strings.ReplaceAll
	fmt.Println(strings.ReplaceAll("Ahmad Ahmad Ahmad", "Ahmad", "Budi"))


}