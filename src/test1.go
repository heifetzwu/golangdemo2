package main

import (
	"fmt"
)

var message chan string

func Bot() {
	for {
		msg := <-message
		fmt.Printf("Bot Print:%s\n", msg)
	}

}

func test1() {

	message = make(chan string)

	go Bot()

	message <- "first message"
	fmt.Println("first message send finish")
	message <- "second message"
	fmt.Println("second message send finish")

	fmt.Println("end")
}
