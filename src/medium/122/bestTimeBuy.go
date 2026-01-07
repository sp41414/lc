package bestTimeBuyII

/*
Logic:

[7,1,5,3,6,4]
 x ^

in the for loop, start at 1 and check if the current element is greater than or less than
the previous element. if the current element is greater than, add the difference of the
currentElement - previousElement and add it into maxProfit.

[7,1,5,3,6,4]
   x ^
5 - 1 = 4
maxProfit = 4

[7,1,5,3,6,4]
     x ^
do nothing

[7,1,5,3,6,4]
       x ^
6 - 3 = 3
maxProfit = 4 + 3 = 7

[7,1,5,3,6,4]
         x ^
do nothing

maxProfit = 7, correct.
*/

func MaxProfit(prices []int) int {
	maxProfit := 0
	if len(prices) <= 1 {
		return maxProfit
	}

	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxProfit += prices[i] - prices[i-1]
		}
	}

	return maxProfit
}
