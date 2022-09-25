package codingchallenges

import (
	"bytes"
	"fmt"
	"sort"
	"testing"
)

func GroupAnagrams(strs []string) [][]string {
	// membuat map kosong
	// key: string karakter yang sudah terurut
	// value: slice of string anagram
	groups := make(map[string][]string)

	for _, str := range strs {
		// mengubah string menjadi slice int karakter ASCII
		b := []byte(str)

		// mengurutkan karakter int karakter ASCII
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})

		//membuat key dari karakter ASCII yang sudah terurut
		key := string(b)

		// tambahkan ke groups
		groups[key] = append(groups[key], str)
	}

	// menyimpan value dari groups
	ret := make([][]string, 0, len(groups))
	for _, v := range groups {
		ret = append(ret, v)
	}
	return ret
}

func GroupAnagrams2(strs []string) [][]string {
	groups := map[string][]string{}
	for _, str := range strs {

		// membuat penghitung karakter
		key := map[string]int{}
		for _, char := range "abcdefghijklmnopqrstuvwxyz" {
			key[string(char)] = 0
		}

		for _, char := range str {
			key[string(char)] += 1
		}

		fmt.Println("INI HABIS DITAMBAH", key)
		// mengubah key menjadi string
		b := new(bytes.Buffer)
		for k, v := range key {
			fmt.Println(k, v)
			fmt.Fprintf(b, "%v:%v ", string(k), v)
		}

		groups[b.String()] = append(groups[b.String()], string(str))
	}

	fmt.Printf("%#v\n", groups)
	return [][]string{{"a"}, {"b"}}
}

func TestGroupAnagrams(t *testing.T) {
	res := GroupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(res)
	res = GroupAnagrams2([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(res)

}

func TestUrutanMap(t *testing.T) {
	GroupHashMap := map[string]int{}
	GroupHashMap["a"] = 1
	GroupHashMap["b"] = 1
	GroupHashMap["c"] = 1
	GroupHashMap["d"] = 1
	GroupHashMap["e"] = 1
	GroupHashMap["f"] = 1

	fmt.Printf("SEBELUM: %#v\n", GroupHashMap)
	for k, v := range GroupHashMap {
		fmt.Printf("%v:%v ", k, v)
	}

	GroupHashMap["d"] += 1
	fmt.Printf("Sesudah: %#v\n", GroupHashMap)
	for k, v := range GroupHashMap {
		fmt.Printf("%v:%v ", k, v)
	}
}
