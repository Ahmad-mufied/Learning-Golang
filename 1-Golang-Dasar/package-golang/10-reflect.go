package main

import (
	"fmt"
	"reflect"
)

/*
Package reflect

--	Dalam bahasa pemrograman, biasanya ada fitur Reflection, dimana kita bisa melihat struktur kode
	kita pada saat aplikasi sedang berjalan
--	Hal ini bisa dilakukan di Go-Lang dengan menggunakan package reflect
--	Reflection sangat berguna ketika kita ingin membuat library yang general sehingga mudah digunakan
--	https://golang.org/pkg/reflect/
*/

//* Kode Program StructTag
type Sample struct {
	Name string `required:"true" max:"11"`
}

type ContohLagi struct {
	Name        string `required:"true"`
	Description string `required:"true"`
}

//* Kode Program Validation Library
func IsValid(data interface{}) bool {
	t := reflect.TypeOf(data)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if field.Tag.Get("required") == "true" {
			value := reflect.ValueOf(data).Field(i).Interface()
			if value == "" {
				return false
			}
		}
	}
	return true
}

func main() {
	sample := Sample{"Ahmad"}

	var sampleType reflect.Type = reflect.TypeOf(sample)

	fmt.Println("Ada", sampleType.NumField(), "field")
	fmt.Println("Field Name: ", sampleType.Field(0).Name)
	fmt.Println(sampleType.Field(0).Tag.Get("required"))
	fmt.Println(sampleType.Field(0).Tag.Get("max"))
	fmt.Println(sampleType.Field(0).Tag.Get("min"))

	sample.Name = ""
	fmt.Println(IsValid(sample))

	contoh := ContohLagi{"ahmad", ""}
	fmt.Println(IsValid(contoh))

}
