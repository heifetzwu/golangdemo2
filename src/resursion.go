package main

import "fmt"

func resursion(n int) {
	z := fact(n)
	fmt.Printf("Factorial : %d\n", z)
}

func fact(n int) int {
	if n == 1 {
		return 1
	} else {
		return n * fact(n-1)
	}
}

func resursionpattern(n int) {
	if n == 1 {
		fmt.Println("*")
	} else {
		for i := 1; i <= n; i++ {
			fmt.Print("* ")
		}
		fmt.Println()
		resursionpattern(n - 1)
	}
}
