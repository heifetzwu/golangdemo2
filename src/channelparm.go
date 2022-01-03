package main

import "fmt"

func greet(roc <-chan string) {
	for c := range roc {
		fmt.Println("Hello " + c + "!")
	}
	// fmt.Println("Hello " + <-roc + "!")
}

func greet2(roc <-chan string) {
	for c := range roc {
		fmt.Println("Hello2 " + c + "!")
	}
	// fmt.Println("Hello " + <-roc + "!")
}
func channelparm() {

	testc()

	fmt.Println("main() started")
	c := make(chan string)
	go greet(c)
	go greet2(c)

	c <- "John"
	c <- "Jack"
	c <- "ada"
	c <- "barber"
	c <- "chris"
	c <- "dick"
	fmt.Println("main() stopped")
}
func testc() {
	roc := make(<-chan int)
	soc := make(chan<- int)
	fmt.Printf("Data type of roc is %T \n", roc)
	fmt.Printf("Data type of soc is %T \n", soc)
}
