package ex_7

import "strconv"

func ParseDigitsOfInteger(num int) []int {
	digits := make([]int, 0)
	numStr := strconv.Itoa(num)

	for _, digit := range numStr {
		digits = append(digits, int(digit))
	}

	return digits
}
