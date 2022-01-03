package main

import "github.com/go-martini/martini"

func martinidemo() {
	m := martini.Classic()
	m.Get("/", func() string {
		return "Hello world!"
	})
	m.Run()
}
