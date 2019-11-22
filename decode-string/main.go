package decode_string

import (
	"strconv"
	"strings"
)

func decodeString(str string) string {
	index := strings.Index(str, "[")
	if index == -1 {
		return str
	}
	//找到前缀数字的index
	numStart := index - 1
	for numStart >= 0 && str[numStart] >= '0' && str[numStart] <= '9' {
		numStart--
	}
	times, _ := strconv.Atoi(str[numStart+1 : index])
	//找到匹配的 ]
	var pieceEnd int
	count := 1
	for i := index + 1; i < len(str); i++ {
		if str[i] == '[' {
			count++
		} else if str[i] == ']' {
			count--
		}
		if count == 0 {
			pieceEnd = i
			break
		}
	}
	piece := strings.Repeat(decodeString(str[index+1:pieceEnd]), times)
	if numStart+1 > 0 {
		pieceLeft := str[0 : numStart+1]
		piece = pieceLeft + piece
	}
	if pieceEnd < len(str)-1 {
		pieceRight := decodeString(str[pieceEnd+1:])
		return piece + pieceRight
	}
	return piece

}

//3[ac4[a]]
//3[a3[a3[a]c]]
//3[a]4[b]
//2[abc]ef3[cd]
//3[a]2[b4[F]c]
func decodeString_stack(str string) string {

	var resStack []string
	var countStack []int

	var countStr string
	var piece string
	for i := 0; i < len(str); i++ {
		if str[i] >= '0' && str[i] <= '9' {
			countStr += str[i : i+1]
		} else if str[i] == '[' {
			if piece != "" {
				resStack = append(resStack, piece)
				piece = ""
			}
			if countStr != "" {
				count, _ := strconv.Atoi(countStr)
				countStack = append(countStack, count)
				countStr = ""
			}
			resStack = append(resStack, "[")
		} else if str[i] == ']' {
			//pop stack
			for j := len(resStack) - 1; j >= 0; j-- {
				if resStack[j] == "[" {
					resStack = resStack[:j]
					break
				}
				piece = resStack[j] + piece
			}

			times := countStack[len(countStack)-1]
			countStack = countStack[:len(countStack)-1]

			piece = strings.Repeat(piece, times)

			resStack = append(resStack, piece)

			piece = ""

		} else {
			piece += str[i : i+1]
		}
	}

	var result string
	for i := 0; i < len(resStack); i++ {
		result += resStack[i]
	}

	return result + piece

}
