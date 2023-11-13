package main

import (
	"fmt"

	"github.com/grsmv/goweek"
)

func days() {
	week, _ := goweek.NewWeek(2015, 46)
	fmt.Println(week.Days)
	week, _ = goweek.NewWeek(2020, 0)
	fmt.Println(week.Days)

	week, _ = goweek.NewWeek(2021, 1)
	fmt.Println(week.Days)

	week, _ = goweek.NewWeek(2021, 2)
	fmt.Println(week.Days)

	week, _ = goweek.NewWeek(2023, 12)
	fmt.Println(week.Days)
}
