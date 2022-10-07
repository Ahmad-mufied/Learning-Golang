package main

import (
	"1-Golang-Dasar/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
	fmt.Println("")
}