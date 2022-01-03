package main

import "fmt"

type profile struct {
	name       string
	address    string
	email      string
	contactnum string
}

func (p profile) getname() string {
	return p.name
}

func (p profile) getaddress() string {
	return p.address
}

func (p profile) getcontact() map[string]string {
	m := map[string]string{"email": p.email, "contactnum": p.contactnum}
	return m
}

func methoddemo() {
	s := profile{
		name:       "Sherlock",
		address:    "221B Baker Street",
		email:      "Sherlock@holmes.com",
		contactnum: "18811904",
	}

	fmt.Println("s.name:", s.name)

	fmt.Println("Name:", s.getname())
	fmt.Println("Address:", s.getaddress())
	fmt.Println("email:", s.getcontact()["email"])
	fmt.Println("contactnum:", s.getcontact())
}
