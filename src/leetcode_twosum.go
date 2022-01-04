package main

import "fmt"

func twoSum(nums []int, target int) []int {
	fmt.Println("twoSum")
	for i, n := range nums {
		fmt.Println(i, n)
		for j := i + 1; j < len(nums); j++ {
			fmt.Println("j= ", nums[j])
			if target == nums[i]+nums[j] {
				return []int{i, j}
			}
		}

	}
	return []int{0, 0}
}

func twoSum2(nums []int, target int) []int {
	fmt.Println("twoSum2")
	m := make(map[int]int)
	// var m map[int]int
	fmt.Println("#1  m=", m)
	for i, n := range nums {
		fmt.Println("i,n =", i, n)
		_, ok := m[n]
		fmt.Println("ok", ok)
		if ok && m[n] >= 0 {
			fmt.Println("#2  m[n]=", m[n])
			return []int{m[n], i}
		}
		m[target-n] = i
		fmt.Println("#3  m=", m)
	}
	return []int{0, 0}
}
