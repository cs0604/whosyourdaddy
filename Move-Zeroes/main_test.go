package Move_Zeroes

import (
	"log"
	"testing"
)

func Test_moveZeroes(t *testing.T) {
	nums := []int{1, 0, 2, 0, 3, 0, 0, 5, 6}
	moveZeroes(nums)
	log.Println(nums)
}
