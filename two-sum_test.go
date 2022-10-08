package codingchallenges

import (
	"fmt"
	"testing"
)

func TestTwoSumByValue(t *testing.T) {
	nums := []int{1, 2, 3, 5, 7, 2, 7, 9}
	target := 8
	hashMap := map[int]int{}
	result := []int{}

	for _, v := range nums {
		if hashMap[v] != 0 {
			result = append(result, target-v)
			result = append(result, v)
			fmt.Println(result)
			return
		}

		hashMap[target-v] = v
	}

	fmt.Println(result)
}

/*
=== RUN   TestTwoSum
[3 5]
--- PASS: TestTwoSum (0.00s)
PASS
ok      coding-challenges       0.044s
*/

func TestTwoValueByIndex(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	hashMap := map[int]int{} // k:v => index
	result := []int{}

	for i, val := range nums {
		if _, found := hashMap[val]; found { // hashmap[val] exists
			result = append(result, hashMap[val])
			result = append(result, i)
			fmt.Println(result)
			return
		} else {
			hashMap[target-val] = i
		}
	}

	fmt.Println(result)
}

func TestSearchKey(t *testing.T) {
	hashMap := map[string]int{}
	a, found := hashMap["yeye"]
	if found {
		fmt.Println("I FOUND IT")
	} else {
		fmt.Println("CANNOT FOUND")
	}

	fmt.Println(a)
}
