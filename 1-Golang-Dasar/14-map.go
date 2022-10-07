package main

import "fmt"

func main() {

	var person map[string]string = map[string]string{
		"name": "Ahmad",
		"address": "Lampung",
	}

	person["title"] = "Programmer"

	fmt.Println(person)
	fmt.Println(person["name"])
	fmt.Println(person["address"])

	var book map[string]string = make(map[string]string)
	book["tittle"] = "Belajar Go-Lang"
	book["author"] = "Ahmad"
	book["ups"] = "Salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))
}