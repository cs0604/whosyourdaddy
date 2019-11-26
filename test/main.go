package main

import (
	"fmt"
	"strconv"
	"strings"
)

func unquoteCodePoint(s string) (string, error) {
	r, err := strconv.ParseInt(strings.TrimPrefix(s, "\\U"), 16, 32)
	return string(r), err
}
func main() {

	N := 3

	for k := 0; k <= N*2-2; k++ {
		for i := 0; i < N; i++ {
			if i > k {
				break
			}
			for j := 0; j < N; j++ {
				if i+j > k {
					break
				}
				if i+j == k {
					fmt.Print(i, ",", j, "  ")
				}
			}
		}
		fmt.Println()
	}

	fmt.Println("\n\n\n")

	for k := 0; k < N; k++ {

		i := 0
		j := k
		for i < N && j >= 0 {
			fmt.Print(i, ",", j, "  ")
			i++
			j--
		}

		fmt.Println()

	}

	for k := 1; k < N; k++ {
		i := k
		j := N - 1
		for i < N && j >= k {
			fmt.Print(i, ",", j, "  ")
			i++
			j--
		}
		fmt.Println()
	}

}
