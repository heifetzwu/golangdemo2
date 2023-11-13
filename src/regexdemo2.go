package main

import (
	"bytes"
	"fmt"
	"regexp"
)

func regexdemo2() {

	match, _ := regexp.MatchString("p([a-z]+)ch", "peach")

	fmt.Println(match)

	r, _ := regexp.Compile("p([a-z]+)ch")

	fmt.Println(1, r.MatchString("peach"))

	fmt.Println(2, r.FindString("peach punch"))

	fmt.Println(3, r.FindStringIndex("peach punch"))

	fmt.Println(4, r.FindStringSubmatch("peach punch"))

	fmt.Println(5, r.FindStringSubmatchIndex("peach punch"))

	fmt.Println(6, r.FindAllString("peach punch pinch", -1))

	fmt.Println(7.1, r.FindAllStringSubmatch("peach punch pinch", -1))
	fmt.Println(7.2, r.FindStringSubmatch("peach punch pinch"))

	fmt.Println(8, r.FindAllStringSubmatchIndex("peach punch pinch", -1))

	fmt.Println(9, r.FindAllString("peach punch pinch", 2))
	fmt.Println(10, r.FindAllString("peach punch pinch", 3))

	fmt.Println(11, r.Match([]byte("peach")))

	r = regexp.MustCompile("p([a-z]+)ch")
	fmt.Println(12, r)

	fmt.Println(13, r.ReplaceAllString("a peach", "<fruit>"))

	in := []byte("a peach")
	out := r.ReplaceAllFunc(in, bytes.ToUpper)
	fmt.Println(14, string(out))
	// 	fmt.Println("match")

	s1 := "Amazon.com, Inc. (AMZN)NasdaqGS - NasdaqGS Real Time Price. Currency in USDAdd to watchlist3,417.43+8.43 (+0.25%)At close: 4:00PM EDT3,423.00 +5.57 (0.16%)After hours: 6:12PM EDT"
	s2 := "Amazon.com, Inc. (IBM)NasdaqGS - NasdaqGS Real Time Price. Currency in USDAdd to watchlist2,222.22+2.22 (+0.25%)"
	ss := s1 + s2
	r1, _ := regexp.Compile("([\\D]*)([,0-9]*[.]+[0-9]*)")

	fmt.Println("ss = ", ss)
	fmt.Println(15.1, r1.FindString(ss))
	sss := r1.FindAllString(ss, -1)
	for i, s := range sss {
		fmt.Printf("15.%d--%s\n", i, s)

	}

	res := r1.FindStringSubmatch(ss)
	resIdx := r1.FindStringSubmatchIndex(ss)
	names := r1.SubexpNames()
	for i, vv := range res {
		if i != -1 {
			fmt.Println(i, "--", names[i], "--", res[i], "--", resIdx[i], vv)
		}
	}

}
