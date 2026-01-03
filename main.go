package main

import "fmt"

func main() {
	var intSlice []int = []int{1, 2, 3}
	fmt.Printf("IntSlice: %v\n", sumSlice(intSlice))
	var floatSlice []float32 = []float32{1.6, 2.33, 9}
	fmt.Printf("IntSlice: %v\n", sumSlice(floatSlice))
}

func sumSlice[T int | float32 | float64](slice []T) T {
	var sum T
	for _, val := range slice {
		sum += val
	}
	return sum
}
