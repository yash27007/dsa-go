package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Println("------Recursion----------")
	// printNtimes(5,"Yashwanth")
	// fmt.Println(printSumOfN(5))
	// fmt.Println(factorial(5))
	// fmt.Println(checkPalindrome("Yashwanth"))
	// fmt.Println(checkPalindrome("Malayalam"))
}

// 1. Print a name N times using recursion

func printNtimes(n int, name string) {
	if n <= 0 {
		return
	}
	if n == 1 {
		fmt.Println(name)
		return
	}
	fmt.Println(name)
	printNtimes(n-1, name)
}

// 2. Print numbers from current to N

func printNnumbers(current int, n int) {
	if current > n {
		return
	}
	fmt.Println(current)
	printNnumbers(current+1, n)
}

// 3. Print the sum of N numbers recursively

func printSumOfN(n int) int {
	if n <= 0 {
		return 0
	}
	return n + printSumOfN(n-1)
}

// 4. Factorial of a number

func factorial(n int) int {

	if n == 0 || n == 1 {
		return 1
	}

	return n * factorial(n-1)
}

// 5. Check if a string is palindrome or not

/*
In other languages we have char but in go we have runes.
Its just an alias for int32


*/

func checkPalindrome(text string) bool {
	runes := []rune(text) // This converts the string to individual characters aka runes
	left := 0
	right := len(runes) - 1

	for left < right {
		lowerLeft := unicode.ToLower(runes[left])
		lowerRight := unicode.ToLower(runes[right])

		if lowerLeft != lowerRight {
			return false
		}
		left++
		right--

	}
	return true
}
