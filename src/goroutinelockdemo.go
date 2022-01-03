package main

import (
	"fmt"
	"sync"
	"time"
)

func withdraw() {
	mu.Lock()
	balance := money
	time.Sleep(0 * time.Millisecond)
	balance -= 100
	money = balance
	mu.Unlock()
	fmt.Println("After withdrawing $1000, balace: ", money)

	wg.Done()
}

var wg sync.WaitGroup
var mu sync.Mutex
var money int = 800

func goroutinelockdemo() {
	fmt.Println("3~~~We have $1500")
	wg.Add(8)
	go withdraw() // 1 first withdraw
	go withdraw() // 2 second withdraw
	go withdraw() // 3 second withdraw
	go withdraw() // 4 second withdraw
	go withdraw() // 5 second withdraw
	go withdraw() // 6 second withdraw
	go withdraw() // 7 second withdraw
	go withdraw() // 8 second withdraw

	wg.Wait()
}
