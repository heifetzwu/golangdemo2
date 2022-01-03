package main

import "fmt"

func variadic() {
	// 	sum(1, 2, 3)
	n := []int{4, 5, 6}
	sum(n...)
}

func sum(nums ...int) {
	fmt.Println("Nums Received", nums)
	fmt.Println("len=", len(nums))
	total := 0
	for i, num := range nums {
		fmt.Printf("i = %d, num= %d\n", i, num)
		total += num
	}
	fmt.Printf("Total : %d\n", total)
}
