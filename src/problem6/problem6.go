/**
The sum of the squares of the first ten natural numbers is,
1'2 + 2'2 + ... + 10'2 = 385

The square of the sum of the first ten natural numbers is,
(1 + 2 + ... + 10)'2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/
package problem6

func difference(max int) int {
	sums, squares := differenceBetween(0, 0, 1, max)
	return squares - sums
}

func differenceBetween(sumOfSquares, squareOfSums, current, max int) (int, int) {
	if current > max {
		return sumOfSquares, squareOfSums * squareOfSums
	}
	return differenceBetween(sumOfSquares+(current*current), squareOfSums+current, current+1, max)
}
