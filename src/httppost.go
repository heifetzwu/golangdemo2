package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func httppost() {
	http.HandleFunc("/login", httppostlogin)
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func httppostlogin(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println("method:", r.Method)
	if r.Method == "GET" {
		t, _ := template.ParseFiles("./templates/login.gtpl")
		log.Println(t.Execute(w, nil))
	} else {
		// 在這邊放真實的驗證
		fmt.Fprint(w, "username: ", r.Form["username"][0])
		fmt.Fprint(w, ", password: ", r.Form["password"][0])
	}
}