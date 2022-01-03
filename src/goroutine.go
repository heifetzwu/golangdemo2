package main

import (
	"fmt"
	// 	"time"
)

func goroutine(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
	}

}
