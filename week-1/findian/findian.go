package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var text string

	fmt.Println("String to find preffix i, suffix n and contains a:")

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text = strings.ToLower(scanner.Text())

	if text[0] == 'i' && text[len(text)-1] == 'n' && strings.Contains(text, "a") {
		fmt.Println("Found!")
	} else {
		fmt.Println("Not Found!")
	}
}
