package main

import "fmt"

func advslice() {
	s := make([]int, 3)
	for i := 0; i < 3; i++ {
		s[i] = i + 1
	}
	s = append(s, 4, 5, 6, 7, 8, 9)
	fmt.Println(" s=", s)
	fmt.Println(" s[:3]=", s[:3])
	fmt.Println(" s[3:]=", s[3:])
	// fmt.Println(" s[3:]...=", s[3:]...)

	// copy slice
	b := make([]int, len(s))
	copy(b, s)
	fmt.Println("b = ", b)

	// cut slice - 3,4 out
	s = append(s[:2], s[4:]...)
	// 	s = append(s[:2], s[4:])
	fmt.Println("cut(3,4) s = ", s)
	fmt.Println("len of cut(3,4) s = ", len(s))
	// del by index
	s = delbyindex(2, s)
	fmt.Println("Deleted index 2 in s = ", s)

}

func delbyindex(i int, a []int) []int {
	a = append(a[:i], a[i+1:]...)
	b := make([]int, len(a))
	fmt.Println("len(a)=", len(a))
	fmt.Println("len(b)=", len(b))
	copy(b, a)
	return a
}

func delbyindex2(i int, a []int) {
	a = append(a[:i], a[i+1:]...)

}
