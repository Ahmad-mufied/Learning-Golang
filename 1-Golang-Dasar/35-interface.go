package main

import "fmt"

/*
Interface

--	Interface adalah tipe data Abstract, dia tidak memiliki implementasi langsung
--	Sebuah interface berisikan definisi-definisi method
--	Biasanya interface digunakan sebagai kontrak
*/

type HashName interface {
	GetName() string
}

func sayHello(hashName HashName) {
	fmt.Println("Hello", hashName.GetName())
}


/*
Implementasi Interface

--	Setiap tipe data yang sesuai dengan kontrak interface, secara otomatis dianggap sebagai interface tersebut
--	Sehingga kita tidak perlu mengimplementasikan interface secara manual
--	Hal ini agak berbeda dengan bahasa pemrograman lain yang ketika membuat interface, kita harus
	menyebutkan secara eksplisit akan menggunakan interface mana
*/
type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}

type Animal struct {
	Name string
}

func (animal Animal) GetName() string {
	return animal.Name
}

func main() {
	var ahmad Person
	ahmad.Name = "Ahmad"
	sayHello(ahmad)

	var mufied Person
	mufied.Name = "Mufied"
	sayHello(mufied)

	cat := Animal{
		Name: "Push",
	}
	sayHello(cat)
}