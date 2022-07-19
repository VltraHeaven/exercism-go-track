package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) (value int) {
	var cardValues = map[string]int{
		"other": 0,
		"ace":   11,
		"king":  10,
		"queen": 10,
		"jack":  10,
		"ten":   10,
		"nine":  9,
		"eight": 8,
		"seven": 7,
		"six":   6,
		"five":  5,
		"four":  4,
		"three": 3,
		"two":   2,
		"one":   1,
	}
	for t, v := range cardValues {
		switch card {
		case t:
			value = v
		}
	}
	return
}

func FirstTurn(card1, card2, dealerCard string) string {
	switch cardTotal, dealerTotal := ParseCard(card1)+ParseCard(card2), ParseCard(dealerCard); {
	case cardTotal == 22:
		return "P"
	case cardTotal == 21:
		if dealerTotal >= 10 {
			return "S"
		}
		return "W"
	case cardTotal >= 17 && cardTotal < 21:
		return "S"
	case cardTotal >= 12 && cardTotal < 17:
		if dealerTotal >= 7 {
			return "H"
		}
		return "S"
	default:
		return "H"
	}
}
