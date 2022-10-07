package main

import "fmt"

type Man struct {
	Name string
}

//* Tanpa Pointer
func (man Man) Married() {
	man.Name = "Mr. " + man.Name
	fmt.Println("Di method", man.Name)
}

//* Menggunakan Pointer
func (man *Man) Married1() {
	man.Name = "Mr. " + man.Name
	fmt.Println("Di method", man.Name)
}

func main() {
	ahmad := Man{"Ahmad"}
	ahmad.Married()
	fmt.Println(ahmad.Name)

	mufied := Man{"Mufied"}
	mufied.Married1()
	fmt.Println(mufied.Name)
}