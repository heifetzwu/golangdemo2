package main

import (
	"errors"
	"fmt"
)

func Excute() {
	panic("wtf...")
}
func errordemo() {

	//defer 在這邊並沒有另外立章節介紹，在這邊簡單說明，他就是用在，這個func 結束前，你最後需要做什麼
	//ex 有些o pen file的行為 最後需要 close 等等
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	Excute()

	fmt.Println("end")
}

func errordemo2() {
	err := errors.New("errorrrrrrr")
	if err != nil {
		fmt.Print(err)
	}
}
