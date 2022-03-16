package arrays

func SumArray(data [5]int) int {
	var sum int
	for _, v := range data {
		sum += v
	}
	return sum
}

func SumSlice(data []int) int {
	var sum int
	for _, v := range data {
		sum += v
	}
	return sum
}

func SumAll(args ...[]int) []int {
	var sum []int
	for _, arg := range args {
		innerSum := SumSlice(arg)
		sum = append(sum, innerSum)
	}
	return sum
}

func SumAllTails(args ...[]int) []int {
	var tailsSlice [][]int
	for _, arg := range args {
		var tail []int
		if len(arg) > 0 {
			tail = arg[1:]
		}
		tailsSlice = append(tailsSlice, tail)
	}
	return SumAll(tailsSlice...)
}
