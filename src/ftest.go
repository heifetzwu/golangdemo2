package main

import "fmt"
import "time"

func ftest() {

	go func(n int) {
		fmt.Println("hello inline function test", n)

	}(1)
	time.Sleep(time.Second * 1)
}
