/*
Create Date: 2023/08/21
Author: @1chooo (Hugo ChunHo Lin)
Version: v0.0.1
*/

package divisor

func CalculateRemainders(dividend, divisor int) []int {
	remainders := make([]int, 0)

	for i := 1; i <= dividend; i++ {
		remainder := i % divisor
		remainders = append(remainders, remainder)
	}

	return remainders
}
