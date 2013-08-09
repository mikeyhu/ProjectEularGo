/**
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
 */
package problem4

import "testing"

func TestLargestPalindromeFrom2DigitProducts(t *testing.T) {
	testLargestPalindrome(t,99,9009)
}

func TestLargestPalindromeFrom3DigitProducts(t *testing.T) {
	testLargestPalindrome(t,999,906609)
}

func testLargestPalindrome(t *testing.T, max, expected int) {
	actual := largestPalindrome(max)
	if actual != expected {
		t.Errorf("largestPalindrome(%v) = %v, want %v", max, actual, expected)
	}
}

