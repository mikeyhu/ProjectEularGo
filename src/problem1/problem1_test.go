/*
If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.

Find the sum of all the multiples of 3 or 5 below 1000.
*/

package problem1

import "testing"

func TestSumOfNaturalNumbersBelow10(t *testing.T) {
	testSumsOfNaturalNumbers(t, 10, 23)
}

func TestSumOfNaturalNumbersBelow1000(t *testing.T) {
	testSumsOfNaturalNumbers(t, 1000, 233168)
}

func testSumsOfNaturalNumbers(t *testing.T, max, expected int) {
	actual := sumOfNaturalNumbersBelow(max)
	if actual != expected {
		t.Errorf("sumOfNaturalNumbersBelow(%v) = %v, want %v", max, actual, expected)
	}
}
