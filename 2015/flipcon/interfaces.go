package main

import "fmt"

type Shape interface {
	area() float64
}

type Rectangle struct{ h, w float64 }

type Square struct{ s float64 }

func (r Rectangle) area() float64 { return r.h * r.w }

func (sq Square) area() float64 { return sq.s * sq.s }

func main() {
	shapes := []Shape{Rectangle{10, 20}, Square{5}}
	for _, s := range shapes {
		fmt.Printf("%#v: area = %.f\n", s, s.area())
	}
}
