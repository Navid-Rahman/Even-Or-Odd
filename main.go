package main

import "fmt"

// main is the entry point of the Advanced Even/Odd Number Analyzer
func main() {
	fmt.Println("🔢 Advanced Even/Odd Number Analyzer 🔢")
	fmt.Println("=====================================")

	// Get numbers from user input
	numbers := getUserNumbers()

	if len(numbers) == 0 {
		fmt.Println("No numbers provided. Exiting...")
		return
	}

	// Analyze each number
	analyses := analyzeNumbers(numbers)

	// Display detailed analysis
	displayDetailedAnalysis(analyses)

	// Calculate and display statistics
	stats := calculateStatistics(analyses)
	displayStatistics(stats)

	// Performance benchmark
	benchmarkAlgorithms(numbers)

	// Interactive mode
	interactiveMode()
}
