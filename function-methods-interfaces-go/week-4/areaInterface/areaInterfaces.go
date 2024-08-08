package main

import "fmt"

type Shape2D interface {
	Area() float64
	Perimeter() float64
}

type Triangle struct {
	RightSide float64
	LeftSide  float64
	Base      float64
	Height    float64
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

func (t Triangle) Perimeter() float64 {
	return t.Base + t.LeftSide + t.RightSide
}

type Rectangle struct {
	Base   float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Base * r.Height
}

func (t Rectangle) Perimeter() float64 {
	return t.Base + t.Height
}

func fits(s Shape2D, size float64) bool {
	if s.Area() <= size && s.Perimeter() <= size {
		return true
	}
	return false
}

func whatObjectAmI(s Shape2D) {
	_, ok := s.(Rectangle)
	if ok {
		fmt.Println("I am a rectangle")
	}
	_, okT := s.(Triangle)
	if okT {
		fmt.Println("I am a triangle")
	}
}

func whatObjectAmISwitch(s Shape2D) {
	switch s.(type) {
	case Rectangle:
		fmt.Println("I am a rectangle")
	case Triangle:
		fmt.Println("I am a tirangle")
	}
}

func main() {
	t1 := Triangle{Base: 10, Height: 10, LeftSide: 10, RightSide: 10}
	r1 := Rectangle{Base: 10, Height: 10}

	fmt.Println(fits(t1, 30))
	fmt.Println(fits(t1, 100))
	fmt.Println(fits(r1, 100))

	whatObjectAmI(t1)
	whatObjectAmI(r1)
	whatObjectAmISwitch(t1)
	whatObjectAmISwitch(r1)
}
