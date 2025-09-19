package main

import (
	"fmt"
	"time"
)

// benchmarkAlgorithms compares different even/odd detection methods
func benchmarkAlgorithms(numbers []int) {
	fmt.Println("\n⏱️  PERFORMANCE BENCHMARK")
	fmt.Println("=========================")

	// Benchmark different even/odd detection methods
	iterations := 1000000

	// Method 1: Modulo operator
	start := time.Now()
	for i := 0; i < iterations; i++ {
		for _, num := range numbers {
			_ = isEven(num)
		}
	}
	moduloTime := time.Since(start)

	// Method 2: Bitwise AND
	start = time.Now()
	for i := 0; i < iterations; i++ {
		for _, num := range numbers {
			_ = isEvenBitwise(num)
		}
	}
	bitwiseTime := time.Since(start)

	fmt.Printf("Modulo operator (%d iterations): %v\n", iterations, moduloTime)
	fmt.Printf("Bitwise AND (%d iterations): %v\n", iterations, bitwiseTime)

	if moduloTime < bitwiseTime {
		fmt.Println("🏆 Winner: Modulo operator is faster!")
	} else {
		fmt.Println("🏆 Winner: Bitwise AND is faster!")
	}

	// Additional benchmarks
	benchmarkPrimeDetection(numbers)
	benchmarkPerfectSquareDetection(numbers)
}

// benchmarkPrimeDetection benchmarks prime number detection
func benchmarkPrimeDetection(numbers []int) {
	fmt.Println("\n🧮 PRIME DETECTION BENCHMARK")
	fmt.Println("===========================")

	iterations := 10000

	start := time.Now()
	for i := 0; i < iterations; i++ {
		for _, num := range numbers {
			_ = isPrime(num)
		}
	}
	primeTime := time.Since(start)

	fmt.Printf("Prime detection (%d iterations): %v\n", iterations, primeTime)
	fmt.Printf("Average time per number: %v\n", primeTime/time.Duration(len(numbers)*iterations))
}

// benchmarkPerfectSquareDetection benchmarks perfect square detection
func benchmarkPerfectSquareDetection(numbers []int) {
	fmt.Println("\n🔲 PERFECT SQUARE DETECTION BENCHMARK")
	fmt.Println("====================================")

	iterations := 100000

	start := time.Now()
	for i := 0; i < iterations; i++ {
		for _, num := range numbers {
			_ = isPerfectSquare(num)
		}
	}
	squareTime := time.Since(start)

	fmt.Printf("Perfect square detection (%d iterations): %v\n", iterations, squareTime)
	fmt.Printf("Average time per number: %v\n", squareTime/time.Duration(len(numbers)*iterations))
}
