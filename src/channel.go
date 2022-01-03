package main

import (
	"fmt"
	"time"
)

func channel() {
	messages := make(chan string)

	go func() { messages <- "ping" }()
	go func() { messages <- "ping2" }()
	msg := <-messages
	fmt.Println(msg)

	msg = <-messages
	fmt.Println(msg)
	// 	chanmsg := make(chan string)
	// 	chanmsg <- "ping 2"
	// 	msg2 := <-chanmsg
	// 	fmt.Println(msg2)
}

func channel1_2() {
	messages := make(chan string)

	// var messages chan string
	// messages = make(chan string)
	go func() { messages <- "ping" }()
	go func() { messages <- "ping3" }()
	// var msg string

	go func() {
		for msg := range messages {
			fmt.Println("msg = ", msg)
		}
	}()

	time.Sleep(1 * time.Second)
	// msg := <-messages
	// fmt.Println("msg = ", msg)
	// msg = <-messages
	// fmt.Println("msg = ", msg)
	// // for msg := range messages {
	// 	fmt.Printf("msg = %s\n", msg)
	// }

}
