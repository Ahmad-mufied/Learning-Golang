package main

import "fmt"

type HashName interface {
	GetName() string
}

func sayHello(hashName HashName) {
	fmt.Println("Hello", hashName.GetName())
}

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