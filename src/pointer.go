package main

import "fmt"

func pointer1() {
	var a *int
	b := 6

	a = &b
	fmt.Println("Address", a)
	fmt.Println("value", *a)

}

func pointer2() {
	// c := 5

}
