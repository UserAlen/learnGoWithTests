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

func SumAllTails(numToSum ...[]int) []int {
	var result []int

	for _, nums := range numToSum {
		tail := nums[1:]
		result = append(result, Sum(tail))
	}
	return result
}

// My version before looking at answer

/* func SumAllTails(numSlices ...[]int) []int {
	var result []int

	for _, numSlice := range numSlices {
		result = append(result, Sum(removeSliceHead(numSlice)))
	}
	return result
}

func removeSliceHead(nums []int) []int {
	var newNums []int

	for i, num := range nums {
		if i == 0 {
			continue
		}
		newNums = append(newNums, num)
	}
	return newNums
} */
