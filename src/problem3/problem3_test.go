/**
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
 */
package problem3

import "testing"

func TestLargestPrimeFactorOf13195(t *testing.T) {
	testLargestPrimeFactor(t,13195,29)
}

func TestLargestPrimeFactorOf600851475143(t *testing.T) {
	testLargestPrimeFactor(t,600851475143,6857)
}

func testLargestPrimeFactor(t *testing.T, max, expected int) {
	actual := largestPrimeFactor(max)
	if actual != expected {
		t.Errorf("largestPrimeFactor(%v) = %v, want %v", max, actual, expected)
	}
}

