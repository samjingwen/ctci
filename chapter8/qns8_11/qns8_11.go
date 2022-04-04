package qns8_11

func Coins(n int) int {
	denoms := []int{25, 10, 5, 1}

	cache := make([][]int, n+1)
	for i, _ := range cache {
		cache[i] = make([]int, len(denoms))
	}

	var iter func(amount, index int) int
	iter = func(amount, index int) int {
		if cache[amount][index] > 0 {
			return cache[amount][index]
		}
		if index >= len(denoms)-1 {
			return 1
		}

		denomAmount := denoms[index]
		ways := 0
		for i := 0; i*denomAmount <= amount; i++ {
			remaining := amount - i*denomAmount
			ways += iter(remaining, index+1)
		}
		cache[amount][index] = ways
		return ways
	}

	return iter(n, 0)
}
