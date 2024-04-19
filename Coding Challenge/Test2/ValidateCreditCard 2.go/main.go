package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func isValidCreditCardNumber(cardNumber string) bool {
	// Check if the card number starts with 4, 5, or 6
	if match, _ := regexp.MatchString(`^[4-6]\d{3}(-?\d{4}){3}$`, cardNumber); !match {
		return false
	}

	// Remove hyphens from the card number
	cardNumber = regexp.MustCompile(`-`).ReplaceAllString(cardNumber, "")

	// Check if the card number has exactly 16 digits
	if len(cardNumber) != 16 {
		return false
	}

	// Check if the card number consists only of digits
	_, err := strconv.ParseInt(cardNumber, 10, 64)
	if err != nil {
		return false
	}

	// Check for consecutive repeated digits
	for i := 0; i < len(cardNumber)-3; i++ {
		if cardNumber[i] == cardNumber[i+1] && cardNumber[i] == cardNumber[i+2] && cardNumber[i] == cardNumber[i+3] {
			return false
		}
	}

	return true
}

func main() {
	cardNumbers := []string{
		"4253625879615786",
		"4424444424442444",
		"5122-2368-7954-3214",
		"44244x4424442444",
		"0525362587961578",
	}

	for _, cardNumber := range cardNumbers {
		if isValidCreditCardNumber(cardNumber) {
			fmt.Printf("%s is a valid credit card number\n", cardNumber)
		} else {
			fmt.Printf("%s is NOT a valid credit card number\n", cardNumber)
		}
	}
}