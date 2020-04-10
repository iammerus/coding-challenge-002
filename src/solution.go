package main

import (
	"fmt"
)

func main() {
	fmt.Println(isPandigital(1234567891))
}

func isPandigital(number uint64) bool {
	// Create an array with 10 elements initialized to zero
	// We'll use this to track the number of times a digit (0 - 9) appears in the
	// result array
	digits := [10]uint64{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for number > 0 {
		// We get the last digit of the number (it will be in the range 0 - 9)
		index := number % 10

		// We add one to the index (last digit). The reason being. If a number is a pandigital
		// by the end of the iteration, every index's value should be at least one. If there's an index
		// that is still zero. Then that number wasn't in the digit
		digits[index]++

		// Remove that last digit
		number = number / 10
	}

	// Finally. Test the result by iterating over the results array
	// Here, all we want to test if there's a digit that is still zero,
	// since if there is. That means that digits wasn't in the number we got
	// Hence the number we go was not a pandigital
	for i := 0; i < 10; i++ {
		if digits[i] == 0 {
			return false
		}
	}

	return true
}
