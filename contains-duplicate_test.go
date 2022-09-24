package codingchallenges

import (
	"fmt"
	"testing"
)

func ContainsDuplicate(nums []int) bool {
	hashMap := map[int]int{}
	for _, num := range nums {
		_, exist := hashMap[num]
		hashMap[num] = 1
		if exist {
			return true
		}
	}

	// fmt.Printf("%#v\n", hashMap)
	return false
}

func TestContainDuplicate(t *testing.T) {
	res := ContainsDuplicate([]int{3, 4, 2, 5})
	fmt.Println("Result: ", res)
}
