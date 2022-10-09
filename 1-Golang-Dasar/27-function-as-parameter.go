package main

import "fmt"

/*
Function as Parameter

--	Function tidak hanya bisa kita simpan di dalam variable sebagai value
--	Namun juga bisa kita gunakan sebagai parameter untuk function lain


Function Type Declarations

--	Kadang jika function terlalu panjang, agak ribet untuk menuliskannya di dalam parameter
--	Type Declarations juga bisa digunakan untuk membuat alias function, sehingga akan mempermudah kita
	menggunakan function sebagai parameter
*/

type Filter func(string) string

func sayHelloWithFilter(name string, filter func(string) string) {
	nameFiltered := filter(name)
	fmt.Println("Hello", nameFiltered)
}

func sayHelloWithFilter1(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Anjing" {
		return "...."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Ahmad", spamFilter)
	sayHelloWithFilter1("Anjing", spamFilter)
}