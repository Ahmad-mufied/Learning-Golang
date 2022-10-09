package main

import (
	"errors"
	"fmt"
)

/*
error Interface

--	Go-Lang memiliki interface yang digunakan sebagai kontrak untuk membuat error, nama interface
	nya adalah error


Membuat Error

--	Untuk membuat error, kita tidak perlu manual.
--	Go-Lang sudah menyediakan library untuk membuat helper secara mudah, yang terdapat di
	package errors (Package akan kita bahas secara detail di materi tersendiri)
*/


func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh 0")
	} else {
		result := nilai / pembagi
		return result, nil
	}
 }


func main() {
	// var contohError error = errors.New("Ups Error")
	// fmt.Println(contohError.Error())
	hasil, err := Pembagian(100, 0)
	if err == nil {
		fmt.Println("Hasil", hasil)
	} else {
		fmt.Println("Error", err.Error())
	}

}