package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
	// "log"
)

func readfiledemo() {
	// read1()
	read2()
}

func read1() {
	dat, _ := ioutil.ReadFile("dunkerk.txt")
	// if err != nil {
	fmt.Println(dat)
	fmt.Println(string(dat))
	// } else {
	// fmt.Println(err)
	// log.Fatal(err)

	// }

}

func read2() {
	f, _ := os.Open("dunkerk.txt")
	b := make([]byte, 11)
	i := 0
	for {
		i++
		r, _ := f.Read(b)
		if r == 0 {
			break
		}
		fmt.Println(r, i)
		// fmt.Println(b)
		fmt.Println(string(b))

	}

}

func read3() {
	f, _ := os.Open("dunkerk.txt")
	b := make([]byte, 100)

	for {
		r, err := f.Read(b)
		if err != nil && err != io.EOF {
			panic(err)

		}
		if r == 0 {
			break
		}
		fmt.Println(string(b[:r]))
	}
}
