package main

import "fmt"

func main() {
	var x [5]int = [5]int{1, 2, 3, 4, 5}
	var s []int = x[1:3]
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	fmt.Println(s[:cap(s)])

	
}