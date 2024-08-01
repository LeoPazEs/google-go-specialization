package main

import (
	"fmt"
	"math"
)

type MyInt int

// Reciever Type passed by value
func (mi MyInt) Double() int {
	return int(mi * 2)
}

type Point struct {
	x float64
	y float64
}

func (mi Point) DistToOrig() float64 {
	t := math.Pow(mi.x, 2) + math.Pow(mi.y, 2)
	return math.Sqrt(t)
}

func main() {
	v := MyInt(3)
	fmt.Println(v.Double())

	p1 := Point{3, 4}
	fmt.Println(p1.DistToOrig())
}
