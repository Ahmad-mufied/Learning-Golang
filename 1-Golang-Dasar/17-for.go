package main

import "fmt"

func main() {
	counter := 1

	for counter <= 10 {
		fmt.Println("Perulangan ke", counter)
		counter++
	}

	//* For dengan Statement
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	fmt.Println()


	slice := []string{"Ahmad", "Mufied", "Nugroho"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i])
	}

	// * For range Slice

	for i, value := range slice {
		fmt.Println("Index", i, "=", value)
	}

	// * For range Map

	person := make(map[string]string)
	person["name"] = "Ahmad"
	person["tittle"] = "Programmer"

	for key, value := range person {
		fmt.Println(key, "=", value)
	}
}