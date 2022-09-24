package codingchallenges

import (
	"fmt"
	"testing"
)

func isAnagram(s string, t string) bool {
	hashMap := map[string]int{}

	// jika jumlah karakternya tidak sama, pasti bukan anagram
	if len(s) != len(t) {
		return false
	}

	for i, valS := range s {
		_, exists := hashMap[string(valS)]
		if exists {
			hashMap[string(valS)] += 1
		} else {
			hashMap[string(valS)] = 1
		}

		_, exists = hashMap[string(t[i])]
		if exists {
			hashMap[string(t[i])] -= 1
		} else {
			hashMap[string(t[i])] = -1
		}
	}
	// fmt.Printf("%#v\n", hashMap)
	for _, v := range hashMap {
		if v != 0 {
			return false
		}
	}

	return true
}

func TestValidAnagram(t *testing.T) {
	res := isAnagram("nagaram", "anagram")
	fmt.Println(res)
	res = isAnagram("rat", "car")
	fmt.Println(res)

}
