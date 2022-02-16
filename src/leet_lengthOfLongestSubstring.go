package main

import "fmt"

func lengthOfLongestSubstring_demo() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
}

func lengthOfLongestSubstring(s string) int {

	fmt.Println(s)
	for i := 0; i < len(s); i++ {
		fmt.Println(i, string(s[i]))
	}

	m := make(map[byte]int)
	res := 0

	for l, r := 0, 0; r < len(s); r++ {
		if index, ok := m[s[r]]; ok {
			// only update the left pointer if
			// it's behind the last position of the visited character
			l = max(l, index)
		}
		res = max(res, r-l+1)
		m[s[r]] = r + 1
	}
	return res
}
func max(n, m int) int {
	if n > m {
		return n
	}
	return m
}
