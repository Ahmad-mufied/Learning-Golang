package main

import (
	"fmt"
	"regexp"
)

/*
Package regexp

--	Package regexp adalah utilitas di Go-Lang untuk melakukan pencarian regular expression
--	Regular expression di Go-Lang menggunakan library C yang dibuat Google bernama RE2
--	https://github.com/google/re2/wiki/Syntax
--	https://golang.org/pkg/regexp
*/

func main() {
	//* Membuat Regexp
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])o")

	//* Mengecek apakah Regexp match dengan string
	fmt.Println(regex.MatchString("eko"))
	fmt.Println(regex.MatchString("eto"))
	fmt.Println(regex.MatchString("eDo"))

	//* Mencari string yang match dengan maximum jumlah hasil
	search := regex.FindAllString("eko eka edo eto eyo eki", 2)
	fmt.Println(search)
}