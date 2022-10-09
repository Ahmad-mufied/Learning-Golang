package main

import "fmt"

/*
Tipe Data Map

-- Pada Array atau Slice, untuk mengakses data, kita menggunakan index Number dimulai dari 0
-- Map adalah tipe data lain yang berisikan kumpulan data yang sama, namun kita bisa menentukan
   jenis tipe data index yang akan kita gunakan
-- Map adalah tipe data kumpulan key-value (kata kunci - nilai), dimana kata kuncinya
   bersifat unik, tidak boleh sama
-- Berbeda dengan Array dan Slice, jumlah data yang kita masukkan ke dalam Map boleh
   sebanyak-banyaknya, asalkan kata kunci nya berbeda, jika kita gunakan kata kunci sama, maka
   secara otomatis data sebelumnya akan diganti dengan data baru
*/

func main() {

	var person map[string]string = map[string]string{
		"name": "Ahmad",
		"address": "Lampung",
	}

	/*
	Function Map
	
			Operasi							Keterangan
	len(map)						Mendapatkan jumlah data di map
	map[key]						Mengambil data di map dengan key
	map[key] = value				Mengubah data di map dengan key
	make(map[TypeKey]TypeValue)		Membuat map baru
	delete(map, key)				Menghapus data di map dengan key
	*/

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