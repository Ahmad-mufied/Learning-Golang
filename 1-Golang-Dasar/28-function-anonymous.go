package main

import (
	"fmt"
)

/*
Anonymous Function

--	Sebelumnya setiap membuat function, kita akan selalu memberikan sebuah nama pada function tersebut
--	Namun kadang ada kalanya lebih mudah membuat function secara langsung di variable atau parameter tanpa harus
	membuat function terlebih dahulu
--	Hal tersebut dinamakan anonymous function, atau function tanpa nama
*/

type BlackList func(string) bool

func registerUser(name string, blacklist BlackList) {
	if blacklist(name) {
		fmt.Println("You are blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

// func blacklistAdmin(name string) bool {
// 	return name == "admin"
// }
// func blacklistRoot(name string) bool {
// 	return name == "root"
// }

func main() {
	blacklist := func(name string) bool {
		return name == "admin"
	}

	registerUser("admin", blacklist)
	registerUser("Ahmad", blacklist)

	registerUser("root", func(name string) bool {
		return name == "root"
	})
	registerUser("Ahmad", func(name string) bool {
		return name == "root"
	})
}