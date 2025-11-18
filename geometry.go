package main

import "fmt"

type geometry interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

func (r rect) area() float64 { // Receivers Fuctions
	return r.width * r.height
}
func (r rect) perim() float64 {
	return 2*r.width + 2*r.height
}

func measure(g geometry) {
	fmt.Println("Area " + fmt.Sprint(g.area()))
	fmt.Println("perim " + fmt.Sprint(g.perim()))
}
