package main

import "fmt"

type userprofile interface {
	getname() string
	getaddress() string
	getcontact() map[string]string
}

func getuserinfo(u userprofile) {
	fmt.Println("Name :", u.getname())
	fmt.Println("Address: ", u.getaddress())
	fmt.Println("Contact Details:", u.getcontact())
	fmt.Println("Email:", u.getcontact()["email"])
}

func interfacedemo() {

	s := profile{
		name:       "Sherlock",
		address:    "221B Baker Street",
		email:      "Sherlock@holme.com",
		contactnum: "18811904",
	}
	getuserinfo(s)
}
