package phonenumber
import (
	"fmt"
	"strings"
	"unicode"
)


func Number(phoneNumber string) (string, error) {
    
    var digits strings.Builder
    
	for _, r := range phoneNumber {
		if unicode.IsDigit(r) {
			digits.WriteRune(r)
		}
	}

	num := digits.String()

	// Handle country code
	if len(num) == 11 {
		if num[0] != '1' {
			return "", fmt.Errorf("invalid country code")
		}
		num = num[1:]
	}

	if len(num) != 10 {
		return "", fmt.Errorf("incorrect number of digits")
	}

	// Area code must start with 2-9
	if num[0] < '2' || num[0] > '9' {
		return "", fmt.Errorf("invalid area code")
	}

	// Exchange code must start with 2-9
	if num[3] < '2' || num[3] > '9' {
		return "", fmt.Errorf("invalid exchange code")
	}

	return num, nil
}

func AreaCode(phoneNumber string) (string, error) {

    num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return num[:3], nil
}

func Format(phoneNumber string) (string, error) {
    
    num, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf(
		"(%s) %s-%s",
		num[:3],
		num[3:6],
		num[6:],
	), nil
}
