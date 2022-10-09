package main

/*
GOPATH

--	GOPATH adalah environment variable yang berisikan lokasi tempat kita menyimpan project dan
	library Go-Lang
--	GOPATH wajib di buat ketika kita mulai membuat aplikasi lebih dari satu file atau butuh
	menggunakan library dari luar
*/

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