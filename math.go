package main

import "math"

// isPrime checks if a number is prime using optimized trial division
func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	if n == 2 {
		return true
	}
	if n%2 == 0 {
		return false
	}

	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		if n%i == 0 {
			return false
		}
	}
	return true
}

// isPerfectSquare checks if a number is a perfect square
func isPerfectSquare(n int) bool {
	if n < 0 {
		return false
	}
	root := int(math.Sqrt(float64(n)))
	return root*root == n
}

// getDivisors returns all divisors of a number
func getDivisors(n int) []int {
	var divisors []int
	if n == 0 {
		return divisors
	}

	absN := n
	if n < 0 {
		absN = -n
	}

	for i := 1; i <= absN; i++ {
		if absN%i == 0 {
			divisors = append(divisors, i)
		}
	}

	return divisors
}

// factorial calculates the factorial of a number with overflow protection
func factorial(n int) int64 {
	if n < 0 {
		return -1 // Undefined for negative numbers
	}
	if n == 0 || n == 1 {
		return 1
	}

	result := int64(1)
	for i := 2; i <= n; i++ {
		result *= int64(i)
		if result < 0 { // Overflow protection
			return -1
		}
	}
	return result
}

// isEven checks if a number is even using modulo operator
func isEven(n int) bool {
	return n%2 == 0
}

// isEvenBitwise checks if a number is even using bitwise AND
func isEvenBitwise(n int) bool {
	return n&1 == 0
}
