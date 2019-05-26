package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	totalNumbers := float64(len(numbers))
	return sum / totalNumbers
}

func main() {
	arguments := os.Args[1:]
	var numSlice []float64
	for _, argument := range arguments {
		number, err := strconv.ParseFloat(argument, 64)
		if err != nil {
			log.Fatal(err)
		}
		numSlice = append(numSlice, number)
	}
	fmt.Printf("%0.2f\n", average(numSlice...))
}
