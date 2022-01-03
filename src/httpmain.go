package main

import (
	"html/template"
	"log"
	"net/http"
)

var indexHTML = `<html>
<head>
	<meta http-equiv="Content-type" content="text/html; charset=utf-8">
	<title>测试</title>
</head>
<body>
	<form action="/page" method="post">
		用户名：<br>
		<input name="username" type="text"><br>
		请输入文本：<br>
		<textarea name="usertext"></textarea><br>
		<input type="submit" value="提交">
	</form>
</body>
</html>`

func index(w http.ResponseWriter, r *http.Request) {

	// fmt.Fprintf(w, indexHTML)
	path := r.URL.Path
	log.Println(path)

	renderHTML(w, path, "no data")
}

func servefile(w http.ResponseWriter, r *http.Request) {
	log.Println("r.URL.Path=", r.URL.Path)
	http.ServeFile(w, r, "website"+r.URL.Path)
}
func httpmain() {

	// http.HandleFunc("/", index)
	http.HandleFunc("/", servefile)
	err := http.ListenAndServe(":9090", nil)
	if err == nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
	}
}

func renderHTML(w http.ResponseWriter, file string, data interface{}) {
	// t, err := template.New(file).ParseFiles("website" + file)
	t, err := template.ParseFiles("website" + file)
	if err != nil {
		log.Println("file not found ", file)
		return
	}
	log.Println("***", file)
	checkErr(err)
	t.Execute(w, data)
}
