/**
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
 */
package problem4

import "strconv"

func largestPalindrome(max int) (largest int) {
	for outer := 2; outer <= max; outer++ {
		for inner := 1; inner < outer; inner++ {
			if numberToCheck := outer * inner; isPalindrome(numberToCheck) {
				if numberToCheck > largest {
					largest = numberToCheck
				}
			}
		}
	}
	return
}

func isPalindrome(number int) bool {
	numberAsString := strconv.FormatInt(int64(number),10)
	for pos := 0; pos < len(numberAsString)/2; pos++ {
		if numberAsString[pos] != numberAsString[len(numberAsString)-1-pos] {
			return false
		}
	}
	return true
}

