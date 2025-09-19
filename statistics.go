package main

import "math"

// calculateStatistics computes comprehensive statistics from number analyses
func calculateStatistics(analyses []NumberAnalysis) Statistics {
	stats := Statistics{
		Min: math.MaxInt32,
		Max: math.MinInt32,
	}

	for _, analysis := range analyses {
		stats.TotalCount++
		stats.Sum += analysis.Value

		if analysis.Value < stats.Min {
			stats.Min = analysis.Value
		}
		if analysis.Value > stats.Max {
			stats.Max = analysis.Value
		}

		if analysis.IsEven {
			stats.EvenCount++
			stats.EvenSum += analysis.Value
		} else {
			stats.OddCount++
			stats.OddSum += analysis.Value
		}

		if analysis.IsPrime {
			stats.PrimeCount++
		}

		if analysis.IsPerfectSquare {
			stats.PerfectSquareCount++
		}
	}

	if stats.TotalCount > 0 {
		stats.Average = float64(stats.Sum) / float64(stats.TotalCount)
	}

	return stats
}

// getPercentage calculates percentage with one decimal place
func getPercentage(part, total int) float64 {
	if total == 0 {
		return 0.0
	}
	return float64(part) / float64(total) * 100
}
