package main

import (
	"fmt"
	"strconv"
)

func strconvdemo2() {
	v := "42"
	if s, err := strconv.ParseUint(v, 10, 32); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}
	// The parse functions return the widest type (float64, int64, and uint64), but if the size argument specifies a narrower width the result can be converted to that narrower type without data loss:
	if s, err := strconv.ParseUint(v, 10, 64); err == nil {
		fmt.Printf("%T, %v\n", int32(s), int32(s))
	}

	if s, err := strconv.ParseUint(v, 16, 64); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

}
