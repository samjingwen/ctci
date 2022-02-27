package qns1_3

func URLify(url *string) {
	var result []int32
	for _, char := range *url {
		if char == ' ' {
			result = append(result, '%')
			result = append(result, '2')
			result = append(result, '0')
		} else {
			result = append(result, char)
		}
	}
	*url = string(result)
}
