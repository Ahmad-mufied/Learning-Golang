package main

import "fmt"

/*
Struct

--	Struct adalah sebuah template data yang digunakan untuk menggabungkan nol atau lebih tipe data
	lainnya dalam satu kesatuan
--	Struct biasanya representasi data dalam program aplikasi yang kita buat
--	Data di struct disimpan dalam field
--	Sederhananya struct adalah kumpulan dari field
*/

type Customer struct {
	 Name, Adress string
	Age int
}

func sayHi(customer Customer, name string){
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

/*
Struct Method

--	Struct adalah tipe data seperti tipe data lainnya, dia bisa digunakan sebagai parameter untuk function
--	Namun jika kita ingin menambahkan method ke dalam structs, sehingga seakan-akan sebuah struct memiliki function
--	Method adalah function
*/

func (a Customer) sayHu(name string) {
	fmt.Println("Huuu", name, "My Name is", a.Name)
}

func main() {

	/*
	Membuat Data Struct

	--	Struct adalah template data atau prototype data
	--	Struct tidak bisa langsung digunakan
	--	Namun kita bisa membuat data/object dari struct yang telah kita buat
	*/
	var ahmad Customer
	ahmad.Name = "Ahmad"
	ahmad.Adress = "Indonesia"
	ahmad.Age = 30

	sayHi(ahmad, "Joko")
	ahmad.sayHu("Nugroho")

	// fmt.Println(ahmad)
	// fmt.Println(ahmad.Name)
	// fmt.Println(ahmad.Adress)
	// fmt.Println(ahmad.Age)

	/*
	Struct Literals

	-- Sebelumnya kita telah membuat data dari struct, namun sebenarnya ada banyak cara yang bisa kita
		gunakan untuk membuat data dari struct
	*/

	joko := Customer{
		Name: "JOKO",
		Adress: "Jawa",
		Age: 10,
	}
	fmt.Println(joko)

	budi := Customer{"Budi", "Jakarta", 35}
	fmt.Println(budi)

}