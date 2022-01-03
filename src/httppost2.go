package main

import (
	"html/template"
	"log"
	"net/http"
)

type ContactDetails struct {
	Email   string
	Subject string
	Message string
}

func httppost2() {
	http.HandleFunc("/", handler2)
	http.ListenAndServe(":3000", nil)
}

func handler2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	tmpl := template.Must(template.ParseFiles("./templates/httppost2form.gtpl"))
	if r.Method != http.MethodPost {
		tmpl.Execute(w, nil)
		return
	}

	details := ContactDetails{
		// Email:   r.FormValue("email"),
		// Subject: r.FormValue("subject"),
		// Message: r.FormValue("message"),

		Email:   r.Form["email"][0],
		Subject: r.Form["subject"][0],
		Message: r.Form["message"][0],
	}

	// do something with details
	_ = details
	log.Println("Email:" + details.Email)
	log.Println("Subject:" + details.Subject)
	log.Println("Message:" + details.Message)
	type sus struct {
		Success bool
	}
	s := sus{Success: true}
	tmpl.Execute(w, s)
	// tmpl.Execute(w, struct{ Success bool }{true})
}
