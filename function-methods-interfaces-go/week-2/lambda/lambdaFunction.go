package main

import "fmt"

func applyIt(aFunc func(int) int, val int) int {
	return aFunc(val)
}

func main() {
	fmt.Println(applyIt(func(x int) int { return x + 1 }, 2))
}
