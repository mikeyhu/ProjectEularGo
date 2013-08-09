/**
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 */
package problem5

import "testing"

func TestSmallestNumberDivisibleBy1to10(t *testing.T) {
	testSmallestNumber(t,10,2520)
}

func TestSmallestNumberDivisibleBy1to20(t *testing.T) {
	testSmallestNumber(t,20,232792560)
}

func testSmallestNumber(t *testing.T, max, expected int) {
	actual := smallestNumber(max)
	if actual != expected {
		t.Errorf("smallestNumber(%v) = %v, want %v", max, actual, expected)
	}
}

