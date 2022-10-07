package main

import "fmt"

type Customer struct {
	 Name, Adress string
	Age int
}

func sayHi(customer Customer, name string){
	fmt.Println("Hello", name, "My Name is", customer.Name)
}

func (a Customer) sayHu(name string) {
	fmt.Println("Huuu", name, "My Name is", a.Name)
}

func main() {
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

	// joko := Customer{
	// 	Name: "JOKO",
	// 	Adress: "Jawa",
	// 	Age: 10,
	// }
	// fmt.Println(joko)

	// budi := Customer{"Budi", "Jakarta", 35}
	// fmt.Println(budi)

}