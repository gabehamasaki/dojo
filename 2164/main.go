package main

import (
	"fmt"
	"math"
)

// https://judge.beecrowd.com/pt/problems/view/2164
func main() {
	input := 0.0
	fmt.Scan(&input)

	result := fast(input)

	fmt.Printf("%.1f\n", result)
}

func fast(n float64) float64 {
	output := (math.Pow((1+math.Sqrt(5))/2, n) - math.Pow((1-math.Sqrt(5))/2, n)) / math.Sqrt(5)

	return output
}
