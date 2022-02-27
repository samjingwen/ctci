package qns1_1

import (
	"sort"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}

func IsUnique1(str string) bool {
	cache := make(map[int32]bool)
	for _, char := range str {
		_, exists := cache[char]
		if exists {
			return false
		}
		cache[char] = true
	}
	return true
}

func IsUnique2(str string) bool {
	s := sortString(str)
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			return false
		}
	}
	return true
}
