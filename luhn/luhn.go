package luhn

func Valid(card string) bool {
	var sum, digitPosition int
	for i := len(card) - 1; i >= 0; i-- {
		symbol := card[i]
		switch {
		case symbol == ' ':
			continue
		case symbol >= '0' && symbol <= '9':
			digit := int(symbol - '0')
			if digitPosition%2 == 1 {
				digit <<= 1
			}
			if digit > 9 {
				digit -= 9
			}
			sum += digit
			digitPosition++
		default:
			return false
		}
	}
	return digitPosition > 1 && sum%10 == 0
}
