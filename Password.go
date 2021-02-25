/*
 * Password rules:
 * at least 7 letters
 * at least 1 number
 * at least 1 upper case
 * at least 1 special character
 */
package main

import (
	"fmt"
	"strings"
	"unicode"
)

func verifyPassword(password string) error {
	var uppercasePresent bool
	var lowercasePresent bool
	var numberPresent bool
	var specialCharPresent bool

	var passLen int
	var errorString string

	for _, ch := range password {
		switch {
		case unicode.IsNumber(ch):
			numberPresent = true
			passLen++
		case unicode.IsUpper(ch):
			uppercasePresent = true
			passLen++
		case unicode.IsLower(ch):
			lowercasePresent = true
			passLen++
		case unicode.IsPunct(ch) || unicode.IsSymbol(ch):
			specialCharPresent = true
			passLen++
		case ch == ' ':
			passLen++
		}
	}
	appendError := func(err string) {
		if len(strings.TrimSpace(errorString)) != 0 {
			errorString += "" + err
		} else {
			errorString = err
		}
	}

	if !lowercasePresent {
		appendError("lowercase letter missing\n")
	}
	if !uppercasePresent {
		appendError("uppercase letter missing\n")
	}
	if !numberPresent {
		appendError("atleast one numeric character required\n")
	}
	if !specialCharPresent {
		appendError("special character missing\n")
	}
	if !(8 <= passLen && passLen <= 64) {
		appendError(fmt.Sprintf("password length must be between 8 to 64 characters long\n"))
	}

	if len(errorString) != 0 {
		return fmt.Errorf(errorString)
	}
	return nil
}

// Let's test it
func main() {
	fmt.Println("enter a password-")
	var password string
	fmt.Scanln(&password)

	err := verifyPassword(password)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("password is set to: %s", password)
	}

}
