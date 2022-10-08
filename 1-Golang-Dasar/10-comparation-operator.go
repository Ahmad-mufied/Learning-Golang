package main

import "fmt"

/*
Operasi Perbandingan

-- operasi untuk membandingkan dua buah data
-- operasi yang menghasilkan nilai boolean (benar atau salah)
-- benar = true
-- salah = false


Operator Perbandingan

Operator			Keterangan
>					Lebih dari
<					Kurang dari
>=					Lebih dari sama dengan
<=					Kurang dari sama dengan
==					sama dengan
!=					tidak sama dengan

*/

func main() {
	
	var name1 = "Ahmad"
	var name2 = "Budi"

	var result bool = name1 == name2
	fmt.Println(result)

	var value1 = 100
	var value2 = 200
	fmt.Println()
	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)

}