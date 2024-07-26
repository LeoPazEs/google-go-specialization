package main

import "os"

func main() {
	var data string = "Hello World"
	err := os.WriteFile("outfile.txt", []byte(data), 0777)
	_ = err
}
