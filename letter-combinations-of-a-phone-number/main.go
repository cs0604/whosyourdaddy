package letter_combinations_of_a_phone_number

import "strconv"

var base = []string{
	"abc",
	"def",
	"ghi",
	"jkl",
	"mno",
	"pqrs",
	"tuv",
	"wxyz",
}

func letterCombinations(digits string) []string {

	var result []string

	if len(digits) > 0 {
		helper("", digits, 0, len(digits), &result)
	}
	return result
}

func helper(cur string, digits string, i int, targetI int, result *[]string) {
	if len(cur) == targetI {
		*result = append(*result, cur)
		return
	}
	index, _ := strconv.Atoi(digits[i : i+1])
	targetJ := len(base[index-2])
	for k := 0; k < targetJ; k++ {
		tmp := cur + base[index-2][k:k+1]
		helper(tmp, digits, i+1, targetI, result)
	}
}
