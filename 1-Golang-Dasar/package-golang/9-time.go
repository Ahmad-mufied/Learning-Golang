package main

import (
	"fmt"
	"time"
)

/*
Package time

--	Package time adalah package yang berisikan fungsionalitas untuk management waktu di Go-Lang
--	https://golang.org/pkg/time/
*/


func main() {
	now := time.Now()
	
	fmt.Println(now)
	fmt.Println(now.Year())
	fmt.Println(now.Month())
	fmt.Println(now.Day())
	fmt.Println(now.Hour())
	fmt.Println(now.Minute())
	fmt.Println(now.Second())
	fmt.Println(now.Nanosecond())

	utc := time.Date(2022, 10, 3, 11, 20, 10, 10, time.UTC)
	fmt.Println(utc)

	layout := "2006-01-02"
	parse, _ := time.Parse(layout, "2022-10-03")
	fmt.Println(parse)
}