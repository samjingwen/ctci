package qns1_2

func CheckPermutation(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}
	s1CharCountMap := getCharCount(s1)
	s2CharCountMap := getCharCount(s2)

	for k, s1Count := range s1CharCountMap {
		s2Count, exists := s2CharCountMap[k]
		if s1Count > s2Count {
			return false
		} else if !exists {
			return false
		}
	}
	return true
}

func getCharCount(s string) map[int32]int {
	charCount := make(map[int32]int)
	for _, val := range s {
		_, exists := charCount[val]
		if exists {
			charCount[val] += 1
		} else {
			charCount[val] = 1
		}
	}
	return charCount
}
