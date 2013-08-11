/**
Find the greatest product of five consecutive digits in the 1000-digit number.
*/
package problem8

import "strconv"

func consecutive(source string, cons int) int {
	largestProduct := 0
	for nextSlide := sliding(stringAsInts(source), cons); ; {
		if next := nextSlide(); next != nil {
			if nextProduct := products(next); nextProduct > largestProduct {
				largestProduct = nextProduct
			}
		} else {
			return largestProduct
		}
	}
}

func products(numbers []int) int {
	product := 1
	for _, p := range numbers {
		product *= p
	}
	return product
}

func stringAsInts(source string) []int {
	sourceInts := make([]int, len(source))
	for p, r := range source {
		sourceInts[p], _ = strconv.Atoi(string(r))
	}
	return sourceInts
}

func sliding(sourceInts []int, cons int) func() []int {
	position := -1
	return func() (next []int) {
		if position++; position+cons > len(sourceInts) {
			return nil
		}
		return sourceInts[position : position+cons]
	}
}
