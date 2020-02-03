package daily_temperatures

//Each temperature will be an integer in the range [30, 100]
func dailyTemperatures2(T []int) []int {

	var result = make([]int, len(T))

	var tbase = make([][]int, 71)

	for i := 0; i < len(T); i++ {
		index := T[i] - 30
		//把小于index的tbase清空，同时产生result
		for j := 0; j < index; j++ {
			arr := tbase[j]
			for k := 0; k < len(arr); k++ {
				result[arr[k]] = i - arr[k]
			}
			tbase[j] = nil
		}
		tbase[index] = append(tbase[index], i)

	}

	return result

}

func dailyTemperatures3(T []int) []int {

	var result = make([]int, len(T))
	if len(T) == 0 {
		return result
	}

	var tbase []int

	tbase = append(tbase, 0)

	for i := 1; i < len(T); i++ {
		if T[i] > T[tbase[0]] {
			for j := 0; j < len(tbase); j++ {
				result[tbase[j]] = i - tbase[j]
			}
			tbase[0] = i
			tbase = tbase[:1]
		} else {
			low := 0
			high := len(tbase) - 1
			for low <= high {
				mid := (low + high) >> 1

				if T[tbase[mid]] >= T[i] {
					low = mid + 1
				} else {
					high = mid - 1
				}
			}

			if low > len(tbase)-1 {
				tbase = append(tbase, i)
			} else {
				for j := low; j < len(tbase); j++ {
					result[tbase[j]] = i - tbase[j]
				}
				tbase[low] = i
				tbase = tbase[0 : low+1]
			}
		}
	}

	return result

}

func dailyTemperatures(T []int) []int {
	var result = make([]int, len(T))
	var stack = make([]int, len(T))
	j := -1
	for i := 0; i < len(T); i++ {
		for j > -1 && T[stack[j]] < T[i] {
			result[stack[j]] = i - stack[j]
			j--
		}
		j++
		stack[j] = i
	}
	return result
}
