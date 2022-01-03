package main

import "fmt"

type Book struct {
	name   string
	author string
	pages  int
}

type Shelf struct {
	position int
	book     Book
}

//點
type Point struct {
    X, Y int
}
//圓
type Circle struct {
    Point
    Radius int
}
//車輪
type Wheel struct {
    Circle
    Spokes int
}
func structdemo() {
	book := Book{name: "Harry Potter", author: "J.K. Rowling", pages: 800}
	fmt.Println(book)
	fmt.Println("Name :", book.name)
	fmt.Println("Author :", book.author)
	fmt.Println("Pages :", book.pages)
	fmt.Println()

	fmt.Println("Book Shelf")
	s := Shelf{position: 1, book: book}
	fmt.Println(s)
	fmt.Println("Book Position:", s.position)
	fmt.Println("Book Details:", s.book)

	var w1 Wheel
	// var w2 Wheel
	w1.Circle.Point.X = 8
	w1.Circle.Point.Y = 8
	w1.Circle.Radius = 5
	w1.Spokes = 20

 	var w2 Wheel
 	w2.X = 11
 	w2.Y = 10

	// w2.X = 8            // 等同於 w.Circle.Point.X = 8
	// w2.Y = 8            // 等同於 w.Circle.Point.Y = 8
	// w2.Radius = 5       // 等同於 w.Circle.Radius = 5
	// w2.Spokes = 20
	
	w3 := Wheel{
	    Circle: Circle{
	        Point:  Point{X: 8, Y: 8},
	        Radius: 5,
	    },
	    Spokes: 20, //特別注意，最後一個成員結尾的逗點是必要的
	}
	// var w3 Wheel
	// w3 = Wheel{
	//      Circle{
	//         Point{X: 8, Y: 8},
	//         Radius: 5,
	//     },
	//     Spokes: 20, //特別注意，最後一個成員結尾的逗點是必要的
	// }
	w4 := Wheel{Circle{Point{8, 8}, 5}, 20}

	fmt.Printf("%v\n", w1)
	fmt.Printf("%+v\n", w2)
	fmt.Printf("%#v\n", w3)
	fmt.Printf("%#v\n", w4)
}
