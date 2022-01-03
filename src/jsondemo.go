package main

import (
	"encoding/json"
	"fmt"
)

// type UserInfo struct {
// 	Name string `json:"name"`
// 	Age  int    `json:"age"`
// }

type UserInfo struct {
	Name string
	Age  int
}

func jsondemo() {
	var jsonString string
	jsonString = `[{"name":"syhlion","age":5} ,{"name":"jack","age":1}]`
	fmt.Println(jsonString)
	//把 json unmarshal 進去 struct
	// u := &UserInfo{}
	var uu []UserInfo
	// u := []UserInfo{}
	err := json.Unmarshal([]byte(jsonString), &uu)
	// u := *uu

	fmt.Println("uu =", uu)
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Printf("name:%s, age:%d\n", u.Name, u.Age)

	//把 struct 轉成 json 字串
	// b, err := json.Marshal(u)
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(string(b))

}
