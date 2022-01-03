package main

import (
	"fmt"
)

func slice_other() {
	var a []string
	a = make([]string, 5, 5)
	a[0] = "A"
	a[1] = "B"
	a[2] = "C"
	a[3] = "D"
	a[4] = "E"
	fmt.Printf("a array %#v\n", a)

	s := a[2:4]
	//s[0] == "C"
	//s[1] == "D"
	//len(s) == 2
	//cap(s) == 3
	s[0] = "Z"
	// a[2] == "Z" 這段是拿來證明 s是參照到 a
	fmt.Printf("s array %#v\n", s)
	fmt.Printf("a array %#v\n", a)
}