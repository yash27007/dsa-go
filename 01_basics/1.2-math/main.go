/*
This file contains solutions to the basic math problems in Striver sheet
*/
package main

import (
	"fmt"
	"math"
)

func main(){
	fmt.Println("Basic Math using go .......")
	fmt.Println(armstrongNumber(153))

}

// 1. Count the number of digits in a number

func countDigits(n float64) int{
	count := int(math.Log10(n) + 1)
	return count
}

func reverseNumber(n int) int{
	rev := 0
	for n > 0 {
		digit := n % 10
		rev = (rev * 10) + digit
		n = n/10
	}
	return rev
}

// 3. Check if a given number is palindrome

func checkPalindrome(n int) bool{
	return reverseNumber(n) == n
}

// 4. GCD of two numbers - Euclidean Algorithm

func gcd(a int, b int) int{
	for a > 0 && b > 0{
		if(a>b){
			a = a % b
		}else{
			b = b % a
		}
	}
	if(a == 0){
		return b
	}
	return a
}

// 5. Armstrong Number

func armstrongNumber(n int) bool{
	dup := n
	sum := 0
	dig := countDigits(float64(n))
	for n > 0{
		ld := n % 10
		sum = sum + int(math.Pow(float64(ld), float64(dig)))
		n = n / 10
	}
	return sum == dup
}