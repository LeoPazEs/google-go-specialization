package data

import "fmt"

var x int = 1

func PrintX() {
	fmt.Println(x)
}

type Point struct {
	x float64
	y float64
}

func (p *Point) InitMe(xn, yn float64) {
	p.x = xn
	p.y = yn
}

func (p *Point) Scale(v float64) {
	p.x *= v
	p.y *= v
}

func (p Point) PrintMe() {
	fmt.Println(p.x, p.y)
}
