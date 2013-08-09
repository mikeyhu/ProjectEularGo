/**
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
 */
package problem5

func smallestNumber(max int) int {
	isDivisible := numIsDivisible(max)
	num := 1
	for ; !isDivisible(num); num++ {}
	return num
}

func numIsDivisible(max int) func(int) bool {
	return func(num int) bool {
		for d := max; d >= 2;d-- {
			if num%d!=0 {
				return false
			}
		}
		return true
	}
}



