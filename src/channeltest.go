package main

import "fmt"

// import "time"

func channeltest() {
	g := make(chan int)
	quit := make(chan bool)

	// 	go B()

	testint := 1
	switch testint {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	case 0:
		fmt.Println("case 0")
	}
	// go func() {
	// 	for {

	// 		select {
	// 		case i := <-g:
	// 			// fmt.Println("go func", <-g)
	// 			fmt.Println("go func", i+1)
	// 		case <-quit:
	// 			fmt.Println("B quit")
	// 			return
	// 		}
	// 	}
	// }()

	go readchan(g, quit)

	for i := 0; i < 5; i++ {
		g <- i
	}
	quit <- true // 沒辦法等待B的退出只能Sleep
	// quit <- false // 沒辦法等待B的退出只能Sleep
	// 	time.Sleep(time.Second * 1)
	fmt.Println("Main quit")
}

func readchan(gmsg chan int, quitflag chan bool) {
	for {

		select {
		case i := <-gmsg:
			// fmt.Println("go func", <-g)
			fmt.Println("go func", i+1)
		case <-quitflag:
			fmt.Println("B quit")
			return
		}
	}
}

// func B() {
// 	for {
// 		select {
// 		case i := <-g:
// 			fmt.Println(i + 1)
// 		case <-quit:
// 			fmt.Println("B quit")
// 			return
// 		}
// 	}
// }
