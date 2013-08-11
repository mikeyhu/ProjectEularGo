/**
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
 */
package problem7

func primeAtPosition(desiredPosition int) int {
	isPrime := isPrimeGen()
	for pos,current := 1,3;;current++ {
		if isPrime(current) {
			if pos++;pos == desiredPosition {
				return current
			}
		}
	}
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

