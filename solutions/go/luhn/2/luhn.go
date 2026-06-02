package luhn
import "strings"

func Valid(id string) bool {
	// Trimmed the string first to get the accurate length without spaces
	input := strings.ReplaceAll(id, " ", "")

	// Valid string length check must happen AFTER removing spaces
	if len(input) <= 1 {
		return false
	}

	sum := 0
	double := false

	for i := len(input) - 1; i >= 0; i-- {
		r := input[i]

		if r < '0' || r > '9' {
			return false
		}

		digit := int(r - '0')

		if double {
			digit *= 2
			if digit > 9 {
				digit -= 9
			}
		}

		sum += digit
		double = !double
	}

	return sum%10 == 0
}