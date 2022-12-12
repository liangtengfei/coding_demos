package utils

import (
	"fmt"
	"net/mail"
	"regexp"
)

var (
	isValidUsername = regexp.MustCompile(`^[a-z0-9_]+$`).MatchString
	isValidFullName = regexp.MustCompile(`^[a-zA-Z\s]+$`).MatchString
)

func ValidateStringLength(value string, minLength int, maxLength int) error {
	n := len(value)
	if n < minLength || n > maxLength {
		return fmt.Errorf("长度必在【%d-%d】之间。", minLength, maxLength)
	}
	return nil
}

func ValidateUsername(value string) error {
	if err := ValidateStringLength(value, 3, 100); err != nil {
		return err
	}
	if !isValidUsername(value) {
		return fmt.Errorf("只能包含小写字母、数字、下划线。")
	}
	return nil
}

func ValidateFullName(value string) error {
	if err := ValidateStringLength(value, 3, 100); err != nil {
		return err
	}
	if !isValidFullName(value) {
		return fmt.Errorf("只能包含字母和空格")
	}
	return nil
}

func ValidatePassword(value string) error {
	return ValidateStringLength(value, 6, 100)
}

func ValidateEmail(value string) error {
	if err := ValidateStringLength(value, 3, 200); err != nil {
		return err
	}
	if _, err := mail.ParseAddress(value); err != nil {
		return fmt.Errorf("无效的邮箱地址。")
	}
	return nil
}
