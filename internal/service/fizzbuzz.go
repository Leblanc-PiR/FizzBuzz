package service

import (
	"strconv"
)

// FizzBuzz handles the actual logic of checking multiples and returning the awaited string
func FizzBuzz(int1, int2, lim int, str1, str2 string) string {
	res := ""

	for i := 1; i <= lim; i++ {
		isMultipleOfInt1 := IsIntModOfIteration(int1, i)
		isMultipleOfInt2 := IsIntModOfIteration(int2, i)

		if !isMultipleOfInt1 && !isMultipleOfInt2 {
			res += strconv.Itoa(i)
		} else {
			if isMultipleOfInt1 {
				res += str1
			}
			if isMultipleOfInt2 {
				res += str2
			}
		}

		if i != lim {
			res += ","
		} else {
			res += "."
		}
	}

	res += "\n"

	return res
}

// IsIntModOfIteration checks if checkedVal is a multiple of iteration, if so returns true and awaited string
func IsIntModOfIteration(checkedVal, iteration int) bool {
	return iteration%checkedVal == 0
}
