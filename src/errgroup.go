package main

import (
	"fmt"

	"golang.org/x/sync/errgroup"
)

func errgroup_demo() {
	var g errgroup.Group
	mp := make(map[string]string)
	g.Go(func() error {
		mp["A"] = "BAD"
		return nil
	})
	g.Go(func() error {
		mp["B"] = "GOOD"
		return nil
	})
	if err := g.Wait(); err != nil {
		fmt.Println("goroutine err:", err)
	}
}

//Found 1 data race(s
