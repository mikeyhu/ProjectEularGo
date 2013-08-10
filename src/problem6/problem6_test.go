/**
The sum of the squares of the first ten natural numbers is,
1'2 + 2'2 + ... + 10'2 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)'2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/
package problem6

import "testing"

func TestDifferenceOfSumsUpTo10(t *testing.T) {
	testDifferenceOfSums(t, 10, 2640)
}

func TestDifferenceOfSumsUpTo100(t *testing.T) {
	testDifferenceOfSums(t, 100, 25164150)
}

func testDifferenceOfSums(t *testing.T, max, expected int) {
	actual := difference(max)
	if actual != expected {
		t.Errorf("difference(%v) = %v, want %v", max, actual, expected)
	}
}
