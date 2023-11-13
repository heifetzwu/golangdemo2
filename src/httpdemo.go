package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"

	// "io/ioutil"

	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gorilla/mux"
)

func httpGet() {
	resp, err := http.Get("https://tw.yahoo.com")

	if err != nil {

	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {

	}

	fmt.Println(string(body))
}

func httpPost() {
	resp, err := http.Post("https://tw.yahoo.com/",
		"application/x-www-form-urlencoded",
		strings.NewReader("name=test"))
	if err != nil {
		fmt.Println("## err = ", err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err = ", err)
	}

	fmt.Println(string(body))
}

func httpClient() {

	client := &http.Client{}

	req, err := http.NewRequest("POST", "https://tw.yahoo.com/", strings.NewReader("name=test"))
	if err != nil {
		// handle error
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Cookie", "name=test")

	resp, err := client.Do(req)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	fmt.Println(string(body))
}

func httpPostForm() {
	resp, err := http.PostForm("https://tw.yahoo.com/",
		url.Values{"key": {"Value"}, "id": {"test"}})
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}

	fmt.Println(string(body))
}

func httpListenAndServe() {
	http.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func httpHandleJson() {
	type UserInfo struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	http.HandleFunc("/api/query", func(w http.ResponseWriter, r *http.Request) {
		u := &UserInfo{
			Name: "syhlion",
			Age:  18,
		}
		b, err := json.Marshal(u)
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(string(b))
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
	// 	http.ListenAndServe(":8080", nil)
}

func muxdemo() {
	r := mux.NewRouter()
	r.HandleFunc("/api/query/{name}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		//這邊就把 name 帶入 UserInfo
		u := &UserInfo{
			Name: vars["name"],
			Age:  18,
		}
		b, err := json.Marshal(u)
		if err != nil {
			log.Println(err)
			return
		}
		w.Header().Set("Content-Type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		w.Write(b)

	})

	//聽8000 port
	log.Fatal(http.ListenAndServe(":8000", r))
}

func httpsitemap() {
	var (
		buf    bytes.Buffer
		logger = log.New(&buf, "logger: ", log.Lshortfile)
		infof  = func(info string) {
			logger.Output(2, info)
		}
	)
	logger.Print("#1 Hello, log file")
	logger.Print("#2 Hello, log file 2")
	logger.Print("#3 Hello, log file 3")

	infof("infof  #")
	// fmt.Println("#4----- &buf=\n", &buf)
	fmt.Println("print something", &buf)
	client := &http.Client{}

	//這邊可以任意變換 http method  GET、POST、PUT、DELETE
	req, err := http.NewRequest("POST", "https://blog.syhlion.tw/sitemap.xml", nil)
	if err != nil {
		fmt.Println("------------------")
		log.Println(err)
		return
	}
	req.Header.Add("If-None-Match", `W/"wyzzy"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
		return
	}
	defer resp.Body.Close()
	sitemap, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("result---------------")

	// 	log.Fatal("log test")
	fmt.Printf("%s", sitemap)
}

var countvar int

func httpcount() {
	http.HandleFunc("/", handler) // each request calls handler
	http.HandleFunc("/count/", counter)
	// http.ListenAndServe("localhost:8123", nil)
	log.Fatal(http.ListenAndServe("localhost:8123", nil))
}

// handler echoes the Path component of the request URL r.
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	keys := r.URL.Query()

	fmt.Fprintf(w, "URL.Path = %q\n", keys)
	fmt.Fprintf(w, "URL.Path = %q\n", keys["abc"])

}

func counter(w http.ResponseWriter, r *http.Request) {
	countvar += 1
	fmt.Fprintf(w, "Count %d\n", countvar)
}
