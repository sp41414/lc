package fizzBuzz

import "fmt"

func FizzBuzz(n int) []string {
	if n <= 0 {
		return []string{}
	}

	output := make([]string, n)

	for i := range n {
		num := i + 1

		if num%3 == 0 && num%5 == 0 {
			output[i] = "FizzBuzz"
			continue
		}

		if num%3 == 0 {
			output[i] = "Fizz"
			continue
		}

		if num%5 == 0 {
			output[i] = "Buzz"
			continue
		}

		output[i] = fmt.Sprintf("%d", num)
	}

	return output

}
