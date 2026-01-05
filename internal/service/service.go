package service

import (
	"errors"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

var (
	ErrEmptyInput = errors.New("входная строка пуста")
)

func ConvertString(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", ErrEmptyInput
	}

	if isMorseCode(input) {
		result := morse.ToText(input)
		return result, nil
	}

	result := morse.ToMorse(input)
	return result, nil
}

func isMorseCode(s string) bool {
	return !strings.ContainsFunc(s, func(r rune) bool {
		return r != '.' && r != '-' && r != ' '
	})
}
