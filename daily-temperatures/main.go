package daily_temperatures

//Each temperature will be an integer in the range [30, 100]
func dailyTemperatures(T []int) []int {

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
