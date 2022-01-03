package main

import "fmt"

// import "time"

func channeldemo() {
	ch := make(chan string, 1)
	go func(msg string) {
		ch <- msg
		fmt.Println("inside func 1")
		ch <- msg
		fmt.Println("inside func 2")
	}("hello")

	// go func() {
	// 	ch <- "hello hello"
	// 	fmt.Println("inside func")
	// }()

	fmt.Println("before recv")
	recv := <-ch
	fmt.Println(recv)
	recv = <-ch
	fmt.Println(recv)

	// fmt.Println ("start")
	// // go func (m string) {
	// //     fmt.Println (m)
	// // } ("test")

	// func (){
	//     fmt.Println ("middle")
	// }()
	// time.Sleep (time.Second * 1)
}
