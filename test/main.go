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

	s, err := unquoteCodePoint("\\U1f600")
	if err != nil {
		// handle error
	}
	fmt.Printf("%s\n", s)

	fmt.Printf("%U", '😀')
}
