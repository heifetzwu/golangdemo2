package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

type Todo struct {
	Title string
	Done  bool
}

type TodoPageData struct {
	PageTitle string
	Todos     []Todo
}

func httphtml() {
	tmpl := template.Must(template.ParseFiles("./layout.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := TodoPageData{
			PageTitle: "My TODO list",
			Todos: []Todo{
				{Title: "Task 1", Done: false},
				{Title: "Task 2", Done: true},
				{Title: "Task 3", Done: true},
			},
		}
		tmpl.Execute(w, data)
	})

	http.ListenAndServe(":8000", nil)
}

func ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong2 \n"))
}

func httpfileserver() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/ping", ping)
	if err := http.ListenAndServe(":8000", nil); err != nil {
		panic(err)
	}
}

func apidemo() {
	http.HandleFunc("/bool/", isHappy)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}

func isHappy(w http.ResponseWriter, r *http.Request) {
	a := strings.Split(r.URL.Path, "=")

	c := "true" == "true2"
	fmt.Println(c)

	aa, bb := a[0], a[1] == "true"
	// aa, bb := a[0], a[1]
	fmt.Fprintln(w, "aa=", aa)
	fmt.Fprintln(w, "bb=", bb)
	if bb {
		fmt.Fprint(w, "I am a Happy man")
	} else {
		fmt.Fprint(w, "I am not a Happy man")
	}
}

func apidemoint() {
	http.HandleFunc("/num/", addint)
	if err := http.ListenAndServe(":3000", nil); err != nil {
		panic(err)
	}
}

func addint(w http.ResponseWriter, r *http.Request) {
	a := strings.Split(r.URL.Path, "/")
	fmt.Fprintf(w, "r.URL.Path =%s \n", r.URL.Path)
	fmt.Fprintf(w, "a = %s \n", a)
	fmt.Fprintf(w, "a0= %s \n", a[0])
	fmt.Fprintf(w, "a1= %s \n", a[1])
	fmt.Fprintf(w, "a2= %s \n", a[2])
	_, _, s := a[0], a[1], a[2]
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	} else {
		fmt.Fprintf(w, "n + 100 = %d\n", i+100)
	}
}

func httpquerydemo() {
	http.HandleFunc("/", queryhandler)
	http.ListenAndServe(":3000", nil)
}

func queryhandler(w http.ResponseWriter, r *http.Request) {

	keys, ok := r.URL.Query()["key"]
	// w.Write(keys)

	// fmt.Fprintf(w, "len = %d\n", len(keys))
	log.Println("len=", len(keys))
	if !ok || len(keys[0]) < 1 {

		fmt.Fprint(w, "Url Param 'key' is missing")
		return
	}

	key := keys[0]
	fmt.Fprint(w, "Url Param 'key' is: "+string(key)+"\n")
	// fmt.Fprint(w, "Url Param 'key' is: "+keys)
	fmt.Println(" key = ", keys)
}
