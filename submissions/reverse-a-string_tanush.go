package main

import "fmt"

func reverseString(s string) string {
	runes := []rune(s)					// runes are used in order to handle special characters like emojis correctly
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {	
		runes[i], runes[j] = runes[j], runes[i]		// swap the characters at position i and j
	}
	return string(runes)					// converts the rune back to string type and returns
}

func main() {
	fmt.Println(reverseString("hello"))      // olleh
	fmt.Println(reverseString("Go ❤️"))       // ❤️ oG
}

