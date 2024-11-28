package phonenumber

import (
	"fmt"
	"regexp"
)

func Number(phoneNumber string) (string, error) {
	rg := regexp.MustCompile(`\D`)
	cleanPhoneNumber := rg.ReplaceAllString(phoneNumber, "")

	if len(cleanPhoneNumber) < 10 || len(cleanPhoneNumber) > 11 {
		return "", fmt.Errorf("invalid count of digits")
	}

	if len(cleanPhoneNumber) == 11 {
		if cleanPhoneNumber[0] == '1' {
			cleanPhoneNumber = cleanPhoneNumber[1:]
		} else {
			return "", fmt.Errorf("Invalid country code")
		}
	}

	if cleanPhoneNumber[0] == '1' || cleanPhoneNumber[0] == '0' {
		return "", fmt.Errorf("Invalid area code!")
	}

	if cleanPhoneNumber[3] == '0' || cleanPhoneNumber[3] == '1' {
		return "", fmt.Errorf("Invalid exchange code!")
	}

	return cleanPhoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	area, err := Number(phoneNumber)

	if err != nil {
		return "", err
	}

	return area[:3], nil
}

func Format(phoneNumber string) (string, error) {
	phoneFormat, err := Number(phoneNumber)

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", phoneFormat[:3], phoneFormat[3:6], phoneFormat[6:]), nil
}
