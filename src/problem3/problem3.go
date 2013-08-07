/**
The prime factors of 13195 are 5, 7, 13 and 29.

What is the largest prime factor of the number 600851475143 ?
 */
package problem3

import "math"

func largestPrimeFactor(input int) (result int) {
	isPrime := isPrimeGen()
	previous := 2
	maxToCheck := int(math.Sqrt(float64(input)))
	for current := 3;current < maxToCheck;current++ {
		if input%current==0 {
			for p := previous;p < current;p++ {
				if isPrime(p) {
					previous = current
				}
			}
			if isPrime(current) {
				result = current
			}

		}
	}
	return
}

func isPrimeGen() func(int) bool {
	var primes []int
	primes = append(primes,2)

	return func(current int) bool {
		for _, prime := range primes {
			if current%prime == 0 {return false}
		}
		primes = append(primes,current)
		return true
	}
}


