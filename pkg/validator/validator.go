package validator

func checkLuhn(card int) int {
	sum := 0
	for i := 0; card > 0; i++ {
		cur := card % 10

		if i%2 == 0 {
			cur = cur * 2
			if cur > 9 {
				cur = cur%10 + cur/10
			}
		}

		sum += cur
		card = card / 10
	}

	return sum % 10
}

func IsLuhnValid(card int) bool {
	return (card%10+checkLuhn(card/10))%10 == 0
}

func CalculateLuhn(card int) int {
	sum := checkLuhn(card)

	if sum == 0 {
		return 0
	}

	return 10 - sum
}
