package main

import (
	"1-Golang-Dasar/helper"
	"fmt"
)

func main() {
	helper.SayHello("Ahmad")
	// helper.sayGoodbye("Ahmad") // error
	fmt.Println(helper.Application)
	// fmt.Println(helper.versi)  // error
}