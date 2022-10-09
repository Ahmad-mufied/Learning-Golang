package main

import (
	"fmt"
	"os"
)

/*
Package os

--	Go-Lang telah menyediakan banyak sekali package bawaan, salah satunya adalah package os
--	Package os berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa
	digunakan disemua sistem operasi)
--	https://golang.org/pkg/os/
*/

func main() {
	arg := os.Args
	fmt.Println("Argument : ")
	fmt.Println(arg)

	fmt.Println(arg[1])
	fmt.Println(arg[2])
	fmt.Println(arg[3])

	name, err := os.Hostname()
	if err == nil {
		fmt.Println("Hostname : ", name)
	} else {
		fmt.Println("Error", err.Error())
	}

	username := os.Getenv("APP_USERNAME")
	password := os.Getenv("APP_PASSWORD")

	fmt.Println(username)
	fmt.Println(password)
}