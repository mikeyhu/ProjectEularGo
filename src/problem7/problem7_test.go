/**
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
 */
package problem7

import "testing"

func Test6thPrime(t *testing.T) {
	testPrimes(t, 6, 13)
}

func Test10001thPrime(t *testing.T) {
	testPrimes(t, 10001, 104743)
}

func testPrimes(t *testing.T, max, expected int) {
	actual := primeAtPosition(max)
	if actual != expected {
		t.Errorf("difference(%v) = %v, want %v", max, actual, expected)
	}
}

