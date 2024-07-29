package main

import "fmt"

func applyIt(aFunc func(int) int, val int) int {
	return aFunc(val)
}

func minusTwo(val int) int {
	return val - 2
}

func plusTwo(val int) int {
	return val + 2
}

func main() {
	fmt.Println(applyIt(minusTwo, 2))
	fmt.Println(applyIt(plusTwo, 2))
}
