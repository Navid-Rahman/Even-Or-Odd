package main

// NumberAnalysis represents detailed analysis of a number
type NumberAnalysis struct {
	Value           int
	IsEven          bool
	IsOdd           bool
	IsPrime         bool
	IsPerfectSquare bool
	Divisors        []int
	Factorial       int64
}

// Statistics holds summary statistics
type Statistics struct {
	TotalCount         int
	EvenCount          int
	OddCount           int
	PrimeCount         int
	PerfectSquareCount int
	Sum                int
	Average            float64
	Min                int
	Max                int
	EvenSum            int
	OddSum             int
}
