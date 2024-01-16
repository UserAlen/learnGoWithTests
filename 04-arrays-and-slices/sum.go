package main

func Sum(numbers []int) int {
	var sum int
	for _, number := range numbers {
		sum += number
	}

	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	// 1 - sums := make([]int, len(numbersToSum))
	// 2 - sums := []int{}
	// 3 - var sums []int

	var sums []int

	/*
		1 -
		for i, numbers := range numbersToSum {
			sums[i] = Sum(numbers)
		}

		2, 3 - worrying less about capacity.
		for _, numbers := range numbersToSum {
			sums = append(sums, Sum(numbers))
		}
	*/

	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums
}
