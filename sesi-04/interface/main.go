// SLide Interface & Empty Interface
package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}

type rectangle struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func (r rectangle) area() float64 {
	return r.height * r.width
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (r rectangle) perimeter() float64 {
	return 2 * (r.height * r.width)
}

func (c circle) volume() float64 {
	return 4 / 3 * math.Pi * math.Pow(c.radius, 3)
}

func print(t string, s shape) {
	fmt.Printf("%s area %v\n", t, s.area())
	fmt.Printf("%s perimeter %v\n", t, s.perimeter())
}

func main() {
	var c1 shape = circle{radius: 5}
	var r1 shape = rectangle{width: 10, height: 10}
	var v interface{}

	v = 20
	value, ok := c1.(circle)

	fmt.Printf("Type of c1: %T\n", c1)
	fmt.Printf("Type of r1: %T\n", r1)

	print("Circle", c1)
	print("Rectangle", r1)

	if ok {
		fmt.Printf("Circle value: %v\n", value)
		fmt.Printf("Circle volume: %f\n", value.volume())
	}

	if val, oke := v.(int); oke {
		v = val * 9
	}
	fmt.Println("Value:", v)
}
