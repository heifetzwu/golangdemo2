package main

import (
	"fmt"
	"strconv"
	"strings"
)

func stringsdemo() {
	//字串轉換大小寫
	//https://golang.org/pkg/strings/#ToLower
	//https://golang.org/pkg/strings/#ToUpper

	upperString := "HELLO WORLD"
	lowerString := "hello world"
	fmt.Println("ToUpper:", strings.ToUpper(lowerString))
	fmt.Println("ToLower:", strings.ToLower(upperString))

	//字串切割，可以指定特定分隔符號轉換成 array
	//https://golang.org/pkg/strings/#Split
	splitString := "a,b,c,d,e,f"
	fmt.Println("Split:", strings.Split(splitString, ","))

	//字串陣列 依照特定分隔符，合併成一個字串
	//https://golang.org/pkg/strings/#Join
	joinArray := []string{"a", "b", "c", "d", "e", "f"}
	fmt.Println("Join:", strings.Join(joinArray, ","))
}
func strconvdemo() {
	var int1 int
	var string1 string
	int1 = 5
	string1 = "5"

	//字串轉數字
	ii, err := strconv.Atoi(string1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("to int :& add one", ii+1)

	//數字轉字串
	fmt.Println("to string :", strconv.Itoa(int1))
}
