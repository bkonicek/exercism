package phonenumber

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

var (
	extraRe = regexp.MustCompile(`\D`)
	numRe   = regexp.MustCompile(`^([2-9][0-9]{2}){2}[0-9]{4}$`)
)

func stripExtra(phoneNumber string) (string, error) {
	phoneNumber = extraRe.ReplaceAllString(phoneNumber, "")
	if len(phoneNumber) < 10 || len(phoneNumber) > 11 {
		return "", errors.New("number must be 10 (or 11 with country code) characters long")
	}
	if len(phoneNumber) == 11 {
		first, err := strconv.ParseInt(phoneNumber[0:1], 10, 32)
		if err != nil {
			fmt.Println(first)
			return "", errors.New("first character is not a valid digit")
		}
		if first != 1 {
			return "", errors.New("invalid country code for NANP")
		}
		phoneNumber = phoneNumber[1:]
	}
	return phoneNumber, nil
}

func checkValidDigits(phoneNumber string) error {
	if !numRe.MatchString(phoneNumber) {
		return errors.New("first digit of area code and exchange code must be 2-9")
	}
	return nil
}

func Number(phoneNumber string) (string, error) {
	phoneNumber, err := stripExtra(phoneNumber)
	if err != nil {
		return "", err
	}
	if err := checkValidDigits(phoneNumber); err != nil {
		return "", err
	}
	return phoneNumber, nil
}

func AreaCode(phoneNumber string) (string, error) {
	phoneNumber, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return phoneNumber[0:3], nil
}

func Format(phoneNumber string) (string, error) {
	number, err := Number(phoneNumber)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("(%s) %s-%s", number[:3], number[3:6], number[6:]), nil
}
