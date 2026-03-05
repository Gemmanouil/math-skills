package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	// Check arguments
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . data.txt")
		return
	}

	filename := os.Args[1]

	// Open file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error: cannot open file:", err)
		return
	}
	defer file.Close()

	var numbers []float64
	scanner := bufio.NewScanner(file)

	// Read numbers line by line
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}

		value, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Error: invalid number in file:", line)
			return
		}

		numbers = append(numbers, value)
	}

	if len(numbers) == 0 {
		fmt.Println("Error: file contains no numbers")
		return
	}

	// --- Average ---
	sum := 0.0
	for _, n := range numbers {
		sum += n
	}
	average := sum / float64(len(numbers))

	// --- Median ---
	sort.Float64s(numbers)
	var median float64
	mid := len(numbers) / 2

	if len(numbers)%2 == 0 {
		median = (numbers[mid-1] + numbers[mid]) / 2
	} else {
		median = numbers[mid]
	}

	// --- Variance ---
	var varianceSum float64
	for _, n := range numbers {
		varianceSum += (n - average) * (n - average)
	}
	variance := varianceSum / float64(len(numbers))

	// --- Standard Deviation ---
	stdDev := math.Sqrt(variance)

	// Print results (rounded integers)
	fmt.Printf("Average: %d\n", int(math.Round(average)))
	fmt.Printf("Median: %d\n", int(math.Round(median)))
	fmt.Printf("Variance: %d\n", int(math.Round(variance)))
	fmt.Printf("Standard Deviation: %d\n", int(math.Round(stdDev)))
}
