package arrayslices

func Sum(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbers ...[]int) []int {
	sums := []int{}

	for _, number := range numbers {
		sums = append(sums, Sum(number))
	}
	return sums
}

func SumAllTails(numbers ...[]int) []int {
	sums := []int{}

	for _, number := range numbers {
		if len(number) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(number[1:]))
		}
	}
	return sums
}
