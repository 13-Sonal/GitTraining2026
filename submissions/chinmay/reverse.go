package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	return string(runes)
}

func main() {

	var s string

	fmt.Println("Enter string you want to be reversed:")
	fmt.Scan(&s)
	fmt.Printf("Rversed string is: %s\n", reverseString(s))

	fmt.Println("Hello Go Working fine!")
}
