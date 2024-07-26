package main

import "fmt"

func fooCopy(y int) {
	y += 1
}

func fooReference(y *int) {
	*y += 1
}

func main() {
	x := 2
	fooCopy(x)
	fmt.Println(x)

	fooReference(&x)
	fmt.Println(x)
}
