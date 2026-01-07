package main

import "fmt"

func fact(n int) int {
	if n == 0 || n == 1 {
		return 1
	}

	return n * fact(n-1)
}

func main() {
	var n int
	fmt.Print("Enter Number: ")
	fmt.Scanln(&n)

	factorial := fact(n)
	fmt.Printf("Factorial of %d is %d\n", n, factorial)
}
