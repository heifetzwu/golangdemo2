package main

import "fmt"

func slice1() {
	x := make([]int, 0)
	x = append(x, 1, 2, 3, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14)
	x = append(x, 15, 16, 17)
	fmt.Println(x)
}

func slice2() {
	x := make([]int, 5, 10)
	// fmt.Println("before any modifitions")
	// fmt.Println("x ->", x)
	// fmt.Println("after append")
	// x = append(x, 1, 2, 3, 4, 5, 6)
	// fmt.Println("x ->", x)
	x[1] = 2
	y := x[1:3]

	x[1] = 100
	y[0] = 99
	x = append(x, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17)
	fmt.Println("x ->", x)
	fmt.Println("x[2:5] -> ", x[2:5])
	fmt.Println("x[:5] -> ", x[:5])
	fmt.Println("x[:] -> ", x[:])
	fmt.Println("y[:] -> ", y[:])
}

func slice3() {
	s := make([]string, 3)
	fmt.Println("String-Slice", s)
	s[0] = "a"
	s[2] = "b"
	s = append(s, "c", "d")

	fmt.Println(s)

}
