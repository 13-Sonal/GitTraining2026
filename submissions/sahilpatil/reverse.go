package main

import "strings"

func reverseString(s string) string {
	chars := strings.Split(s, "")
	n := len(chars)

	for i := 0; i < n/2; i++ {
		temp := chars[i]
		chars[i] = chars[n-1-i]
		chars[n-1-i] = temp
	}

	return strings.Join(chars, "")
}
