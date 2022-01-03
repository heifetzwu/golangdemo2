package main

import "fmt"

func defertest() {
	defer func() {
		fmt.Println("first")
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("end")
	}()
	deferf()
}

func deferf() {
	fmt.Println("test")
	panic(2)
	// fmt.Println("test2")
}

func defertest2() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func defertest3main() {
	fmt.Println(defertest3())
}
func defertest3() (i int, j int) {
	defer func() { i++ }()
	defer func() { i++ }()
	return 1, 1
}

type MyError struct {
	Title   string
	Message string
}

func (e MyError) Error() string {
	return fmt.Sprintf("%v: %v", e.Title, e.Message)
}

func errordemo3() {
	err := MyError{"Error title 1", "Error Message 1"}
	fmt.Println(err)

	err = MyError{
		Title:   "Error Title 2",
		Message: "Error Message 2",
	}

	fmt.Println(err)
}
