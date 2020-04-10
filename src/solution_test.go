package main

import (
	"testing"
)

func TestSolution(t *testing.T) {
	tests := map[uint64]bool{
		1023456789:         true,
		1234567890:         true,
		8549176320:         true,
		1234567889:         false,
		120:                false,
		3816547290:         true,
		9814072356:         true,
		8328393134:         false,
		9876543210:         true,
		1224425565:         false,
		98140723568910:     true,
		9086452314890:      false,
		112233445566778899: false,
	}

	for key, element := range tests {
		output := isPandigital(key)

		if element != output {
			t.Errorf("Panidigital test for '%d' was incorrect, expected '%t', got '%t'", key, element, output)
		}
	}
}
