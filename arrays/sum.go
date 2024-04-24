package main

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	lenOfArgs := len (numbersToSum)
	sums := make([]int, lenOfArgs)

	for index, numbers := range numbersToSum {
		sums[index] = Sum(numbers)
	}

	return sums
}