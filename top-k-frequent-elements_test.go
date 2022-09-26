package codingchallenges

import (
	"fmt"
	"sort"
	"testing"
)

func TopKFrequentElements(nums []int, k int) []int {
	hashMap := map[int]int{}

	for _, num := range nums {
		_, exists := hashMap[num]
		if exists {
			hashMap[num] += 1
		} else {
			hashMap[num] = 1
		}
	}

	result := []int{}
	for _, v := range hashMap {
		result = append(result, v)
	}

	sort.Slice(result, func(i, j int) bool {
		return i > j
	})

	return result[:k]
}

func TestTopKFrequentElements(t *testing.T) {

	res := TopKFrequentElements([]int{1, 1, 1, 2, 2, 3}, 2)
	fmt.Printf("%#v\n", res)

}
