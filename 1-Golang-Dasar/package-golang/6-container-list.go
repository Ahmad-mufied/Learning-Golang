package main

import (
	"container/list"
	"fmt"
)



func main() {
	data := list.New()

	data.PushBack("Ahmad")
	data.PushBack("Mufied")
	data.PushBack("Nugroho")
	data.PushFront("Budi")

	//* next value
	fmt.Println(data.Front().Next().Value)

	//* prev value
	fmt.Println(data.Back().Prev().Value)

	fmt.Println("Data Awal =>", data.Front().Value)
	fmt.Println("Data Akhir =>", data.Back().Value)

	fmt.Println("=============")
	//* To iterate over a list 
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}

	fmt.Println("=============")
	//* To reverse list
	for element := data.Back(); element != nil; element = element.Prev() {
		fmt.Println(element.Value)
	}


}