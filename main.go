package main

import "fmt"

func main() {
	currentNumber, previousNumber, sum := 1, 0, 0
	for currentNumber < 4000000 {
		// Check if the current number of the Fibonacci series is even.
		if currentNumber % 2 == 0 {
			// For an even number, add it to the sum
			sum += currentNumber
		}
		// Archive the current number as previousNumber and use the
		// previousNumber to generate the next number in the series
		previousNumber, currentNumber = currentNumber, currentNumber+previousNumber
	}
	fmt.Println(sum)
}
