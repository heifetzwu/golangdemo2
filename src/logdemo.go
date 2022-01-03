package main

import (
	"log"
)

func logdemo() {
	//最基礎的 印出 hello world
	log.Println("[standard] hello world")

	//設定輸出的格式
	log.SetFlags(log.Ltime)
	log.Println("[setflag]  hello world")

	//設定前綴
	log.SetPrefix("[ithome]")
	log.Println("[prefix & setflag] hello world")

}
