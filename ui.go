package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// getUserNumbers prompts user for input and returns a slice of numbers
func getUserNumbers() []int {
	fmt.Println("\nEnter numbers (separated by spaces or one per line, 'done' to finish):")

	scanner := bufio.NewScanner(os.Stdin)
	var numbers []int

	for {
		fmt.Print("Enter number (or 'done'): ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "done" {
			break
		}

		// Try to parse as single number
		if num, err := strconv.Atoi(input); err == nil {
			numbers = append(numbers, num)
			continue
		}

		// Try to parse as space-separated numbers
		parts := strings.Fields(input)
		for _, part := range parts {
			if num, err := strconv.Atoi(part); err == nil {
				numbers = append(numbers, num)
			} else {
				fmt.Printf("⚠️  Invalid number: %s\n", part)
			}
		}
	}

	return numbers
}

// displayDetailedAnalysis shows comprehensive analysis for each number
func displayDetailedAnalysis(analyses []NumberAnalysis) {
	fmt.Println("\n📊 DETAILED ANALYSIS")
	fmt.Println("===================")

	for _, analysis := range analyses {
		fmt.Printf("\n🔢 Number: %d\n", analysis.Value)

		// Even/Odd status with emojis
		if analysis.IsEven {
			fmt.Printf("   📊 Even/Odd: EVEN ✅\n")
		} else {
			fmt.Printf("   📊 Even/Odd: ODD ❌\n")
		}

		// Additional properties
		if analysis.IsPrime {
			fmt.Printf("   🧮 Prime: YES 🟢\n")
		} else {
			fmt.Printf("   🧮 Prime: NO 🔴\n")
		}

		if analysis.IsPerfectSquare {
			fmt.Printf("   🔲 Perfect Square: YES 🟦\n")
		} else {
			fmt.Printf("   🔲 Perfect Square: NO ⬜\n")
		}

		// Divisors
		fmt.Printf("   📋 Divisors: %v\n", analysis.Divisors)

		// Factorial (with overflow protection)
		if analysis.Factorial == -1 {
			fmt.Printf("   🔢 Factorial: Overflow or undefined\n")
		} else {
			fmt.Printf("   🔢 Factorial: %d\n", analysis.Factorial)
		}
	}
}

// displayStatistics shows summary statistics
func displayStatistics(stats Statistics) {
	fmt.Println("\n📈 STATISTICS SUMMARY")
	fmt.Println("=====================")
	fmt.Printf("Total Numbers: %d\n", stats.TotalCount)
	fmt.Printf("Even Numbers: %d (%.1f%%)\n", stats.EvenCount, getPercentage(stats.EvenCount, stats.TotalCount))
	fmt.Printf("Odd Numbers: %d (%.1f%%)\n", stats.OddCount, getPercentage(stats.OddCount, stats.TotalCount))
	fmt.Printf("Prime Numbers: %d (%.1f%%)\n", stats.PrimeCount, getPercentage(stats.PrimeCount, stats.TotalCount))
	fmt.Printf("Perfect Squares: %d (%.1f%%)\n", stats.PerfectSquareCount, getPercentage(stats.PerfectSquareCount, stats.TotalCount))
	fmt.Printf("Sum: %d\n", stats.Sum)
	fmt.Printf("Average: %.2f\n", stats.Average)
	fmt.Printf("Min: %d\n", stats.Min)
	fmt.Printf("Max: %d\n", stats.Max)
	fmt.Printf("Sum of Even: %d\n", stats.EvenSum)
	fmt.Printf("Sum of Odd: %d\n", stats.OddSum)

	// Fun facts
	displayFunFacts(stats)
}

// displayFunFacts shows interesting insights about the data
func displayFunFacts(stats Statistics) {
	fmt.Println("\n🎯 FUN FACTS")
	fmt.Println("===========")
	if stats.EvenCount > stats.OddCount {
		fmt.Println("📊 Even numbers dominate!")
	} else if stats.OddCount > stats.EvenCount {
		fmt.Println("🔢 Odd numbers dominate!")
	} else {
		fmt.Println("⚖️  Perfect balance between even and odd!")
	}

	if stats.PrimeCount > 0 {
		fmt.Printf("🧮 Found %d prime number(s)!\n", stats.PrimeCount)
	}

	if stats.PerfectSquareCount > 0 {
		fmt.Printf("🔲 Found %d perfect square(s)!\n", stats.PerfectSquareCount)
	}
}

// interactiveMode provides real-time number analysis
func interactiveMode() {
	fmt.Println("\n🎮 INTERACTIVE MODE")
	fmt.Println("==================")
	fmt.Println("Enter numbers to analyze in real-time (type 'quit' to exit):")

	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Enter number: ")
		scanner.Scan()
		input := strings.TrimSpace(scanner.Text())

		if strings.ToLower(input) == "quit" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid number. Please try again.")
			continue
		}

		analysis := analyzeSingleNumber(num)
		displayQuickAnalysis(analysis)
	}

	fmt.Println("👋 Thanks for using the Advanced Even/Odd Analyzer!")
}

// displayQuickAnalysis shows a quick summary for a single number
func displayQuickAnalysis(analysis NumberAnalysis) {
	fmt.Printf("🔢 %d is ", analysis.Value)
	if analysis.IsEven {
		fmt.Print("EVEN")
	} else {
		fmt.Print("ODD")
	}

	if analysis.IsPrime {
		fmt.Print(" and PRIME")
	}

	if analysis.IsPerfectSquare {
		fmt.Print(" and a PERFECT SQUARE")
	}

	fmt.Println()
}
