package main

import "fmt"

func whatday(n int) {
	if n != 0 && n <= 7 {
		switch n {
		case 1:
			fmt.Println("Monday")
		case 2:
			fmt.Println("Tuesday")
		case 3:
			fmt.Println("Wednesday")
		case 4:
			fmt.Println("Thursday")
		default:
			fmt.Println("It is weekday")
		}
	} else {
		fmt.Println("wrong input 1-7 accepted")
	}

}
