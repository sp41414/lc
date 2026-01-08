package reversed

import "math"

func Reverse(x int) int {
	reversed := 0

	for x != 0 {
		ld := x % 10
		reversed = reversed*10 + ld
		if reversed > math.MaxInt32 || reversed < math.MinInt32 {
			return 0
		}
		x /= 10
	}

	return reversed
}
