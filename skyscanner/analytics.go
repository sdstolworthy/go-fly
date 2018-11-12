package skyscanner

import (
	"math"
)

func average(nums ...float64) float64 {
	var total float64
	for _, v := range nums {
		total += v
	}

	return total / float64(len(nums))
}

func standardDeviation(nums ...float64) float64 {
	avg := average(nums...)
	var squaredDifference []float64
	for _, v := range nums {
		squaredDifference = append(squaredDifference, math.Pow(v-avg, 2))
	}
	// sigma:
	return math.Sqrt(average(squaredDifference...))
}
