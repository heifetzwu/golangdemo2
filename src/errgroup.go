package main

import (
	"fmt"

	"sync"

	"golang.org/x/sync/errgroup"
)

func errgroup_demo() {
	var g errgroup.Group
	var m sync.RWMutex
	mp := make(map[string]string)
	fmt.Println("errgroup_demo -----------")
	g.Go(func() error {
		m.Lock()
		mp["A"] = "BAD"
		m.Unlock()
		return nil
	})
	g.Go(func() error {
		m.Lock()
		mp["B"] = "GOOD"
		m.Unlock()
		return nil
	})
	if err := g.Wait(); err != nil {
		fmt.Println("goroutine err:", err)
	}
	fmt.Printf("mp = %v\n", mp)
}

func errgroup_demo_datarace() {
	var g errgroup.Group
	// var m sync.Mutex
	mp := make(map[string]string)
	fmt.Println("errgroup_demo_datarace -----------")
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
	fmt.Printf("mp = %v\n", mp)
}

//Found 1 data race(s
