package main

// analyzeNumbers performs comprehensive analysis on a slice of numbers
func analyzeNumbers(numbers []int) []NumberAnalysis {
	var analyses []NumberAnalysis

	for _, num := range numbers {
		analysis := NumberAnalysis{
			Value:           num,
			IsEven:          isEven(num),
			IsOdd:           !isEven(num),
			IsPrime:         isPrime(num),
			IsPerfectSquare: isPerfectSquare(num),
			Divisors:        getDivisors(num),
			Factorial:       factorial(num),
		}
		analyses = append(analyses, analysis)
	}

	return analyses
}

// analyzeSingleNumber analyzes a single number and returns its analysis
func analyzeSingleNumber(num int) NumberAnalysis {
	return NumberAnalysis{
		Value:           num,
		IsEven:          isEven(num),
		IsOdd:           !isEven(num),
		IsPrime:         isPrime(num),
		IsPerfectSquare: isPerfectSquare(num),
		Divisors:        getDivisors(num),
		Factorial:       factorial(num),
	}
}
