package main

import "fmt"

func main() {
	// Reverse A String
	var input1 string
	fmt.Print("Enter your string: ")
	fmt.Scan(&input1)

	reversedString := reverseString(input1)
	fmt.Println("Reversed string is: ", reversedString)

	// Fiboanacci Sequence
	var i int
	fmt.Print("Enter the length of fibonacci sequence: ")
	fmt.Scanln(&i)

	fibonacciSeries := nFibonacci(i)
	fmt.Println(fibonacciSeries)

	// Palindrome Check
	var input2 string
	fmt.Print("Enter your text: ")
	fmt.Scanln(&input2)

	isPalindrome := palindromeCheck(input2)

	if isPalindrome {
		fmt.Printf("%s is a palidrome\n", input2)
	} else {
		fmt.Printf("%s is not a palidrome\n", input2)
	}

	// Factorial
	var j int
	fmt.Print("Enter number: ")
	fmt.Scanln(&j)
	fmt.Printf("Factorial of %d is %d\n", j, factorial(j))
}
