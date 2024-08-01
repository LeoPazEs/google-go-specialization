package main

import "encapsulation/data"

func main() {
	data.PrintX()
	var p data.Point
	p.InitMe(3, 4)
	p.Scale(2)
	p.PrintMe()
}
