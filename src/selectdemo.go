package main

import (
	"fmt"
)

var sender1Channel chan string
var sender2Channel chan string

func Sender1() {
	sender1Channel <- "this is sender1"

}
func Sender2() {
	sender2Channel <- "this is sender2"
}

func selectdemo() {
	sender1Channel = make(chan string)
	sender2Channel = make(chan string)
	go Sender1()
	go Sender2()
	i := 0

	for {

		select {
		case msg := <-sender1Channel:
			fmt.Println(msg)
		case msg := <-sender2Channel:
			fmt.Println(msg)
		}

		//這邊因為只有兩個 sender 各送一條訊息，所以這個loop 只要執行兩次就可以 break，因為不會再有其他訊息進來
		i = i + 1
		if i >= 2 {
			break
		}

	}
	fmt.Println("end")
}
