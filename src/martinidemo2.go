package main

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
)

func martinidemo2() {
	m := martini.Classic()

	m.Get("/hello/:name", func(params martini.Params) string {
		return "Hello " + params["name"]
	})

	m.Use(render.Renderer(render.Options{
		Layout: "layout",
	}))

	type Member struct {
		Id   int
		Name string
		Sex  string
	}
	member := Member{
		Id: 1, Name: "Negaihoshi", Sex: "Male",
	}

	m.Get("/", func(r render.Render) {
		r.HTML(200, "hello2", member)
	})

	m.Run()
}
