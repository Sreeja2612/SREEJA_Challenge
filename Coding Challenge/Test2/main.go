package main

import (
	"fmt"
	"regexp"
	"strings"
)

// ► It must start with a 4,5  or 6. : done
// ► It must contain exactly 16 digits. : done
// ► It must only consist of digits (0-9). : done
// ► It may have digits in groups of 4, separated by one hyphen "-". : done
// ► It must NOT use any other separator like ' ' , '_', etc. : done
// ► It must NOT have 4 or more consecutive repeated digits. : done

func main() {
	sample := []string{
		"4123456789123456",          //valid
		"4123456788883456",          //valid
		"5123-4567-8912-3456",       //valid
		"5123-4444-8912-3456",       //valid
		"4123356789123456",          //valid
		"0123356789123456",          //invalid
		"5894-45267-8912-3456-9087", //invalid
		"4123456789123sss456",       //invalid
		"5894-45267-3456-9087",      //invalid
		"4098-j267-3456-9087",       //invalid
		"61234--8912-3456",          // Invalid, because the card number is not divided into equal groups of .
		"51-67-8912-3456",           // : Invalid, consecutive digits  is repeating  times.
		"5123456789123456",          // : Invalid, because space '  ' and - are used as separators.
		"4444",                      //invalid
		"4444-4444",                 //invalid
		"4444-4444-4444",            //invalid
		"4444-4444-4444-4444",       //invalid
		"4234_4321_5678_8765",       //invalid
		"4234-4321-5678-8765",
		"4234@4321-5678-8765",       //invalid
		"5123 - 3567 - 8912 - 3456", //invalid
		"4123456789123456",          //valid
		"5123-4567-8912-3456",       //valid
		"61234-567-8912-3456",       //invalid
		"4123356789123456",          //valid
	}

	for i := 0; i < len(sample); i++ {
		fmt.Println()
		fmt.Println()
		fmt.Println()
		var validString bool
		if !checkFirstChar(sample[i]) {
			fmt.Println("first char invalid : ", sample[i])
			continue
		}

		sampleSplit := strings.Split(sample[i], "-")

		switch len(sampleSplit) {
		case 1:
			if containsRepeatedChars(string(sampleSplit[0])) {
				fmt.Println("repeated chars in string : ", sampleSplit[0])
				validString = false
				break
			}
			if !checkInvalidChar(string(sampleSplit[0])) {
				fmt.Println("invalid chars in string : ", sampleSplit[0])
				validString = false
				break
			}
			if !checkStrLen(sampleSplit[0], 16) {
				fmt.Println("invalid len : ", sampleSplit[0])
				validString = false
				break
			}
			validString = true
		case 4:
			var finalStr string
			for j := 0; j < len(sampleSplit); j++ {
				if !checkInvalidChar(string(sampleSplit[j])) {
					fmt.Println("invalid chars in string : ", sampleSplit[j])
					finalStr = ""
					validString = false
					break
				}
				if !checkStrLen(string(sampleSplit[j]), 4) {
					fmt.Println("invalid len : ", sampleSplit[j])
					finalStr = ""
					validString = false
					break
				}
				finalStr += string(sampleSplit[j])
			}

			if finalStr != "" {
				validString = true
			}

			if finalStr != "" && containsRepeatedChars(finalStr) {
				fmt.Println("repeated chars in string : ", sample[i])
				validString = false
				continue
			}

		default:
			fmt.Println("invalid string : ", sample[i])
		}
		fmt.Println("=============================================")
		fmt.Println("input : ", sample[i], " valid credit card no. : ", validString)
		fmt.Println("=============================================")
		fmt.Println()
		fmt.Println()
		fmt.Println()
	}
}

func checkFirstChar(input string) bool {
	var status bool
	if string(input[0]) == "4" || string(input[0]) == "5" || string(input[0]) == "6" {
		status = true
	}
	return status
}

func checkStrLen(input string, limit int) bool {
	var status bool
	if len(input) == limit {
		status = true
	}
	return status
}

func checkInvalidChar(input string) bool {
	sampleRegexp := regexp.MustCompile(`^\d+$`)
	return sampleRegexp.MatchString(input)
}

func containsRepeatedChars(s string) bool {
	var lastChar rune
	var lastCharCount = 0
	for _, c := range s {
		if c == lastChar {
			lastCharCount++
			if lastCharCount >= 4 {
				return true
			}
		} else {
			lastChar = c
			lastCharCount = 1
		}
	}

	return false
}

// func validateString(input string) bool {

// 	return ""
// }
