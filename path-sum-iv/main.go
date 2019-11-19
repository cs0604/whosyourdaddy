package path_sum_iv

func getHundreds(num int) int {
	return num / 100
}

func getTens(num int) int {
	return num % 100 / 10
}

func getUnits(num int) int {
	return num % 10
}

func copyAndClearSlice(from, to []int) {
	for i := 0; i < len(from); i++ {
		to[i] = from[i]
		from[i] = 0
	}
}

func pathSum(input []int) int {

	level1 := make([]int, 8)
	level2 := make([]int, 8)

	curLevel := 0

	for i := 0; i < len(input); i++ {
		if getHundreds(input[i]) == 1 {
			level2[0] = getUnits(input[i])
			curLevel = 1
			continue
		}
		if getHundreds(input[i]) > curLevel {
			copyAndClearSlice(level2, level1)
			curLevel++
		}
		pos := getTens(input[i])
		level2[pos-1] = level1[(pos-1)/2] + getUnits(input[i])
	}

	sum := 0

	for i := 0; i < len(level2); i++ {
		sum += level2[i]
	}

	return sum
}
