package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputString string) (string, error) {
	var symbolForRepeat rune

	isPreviousSymbolSlash := false

	var builder strings.Builder

	for _, r := range inputString {
		if unicode.IsDigit(r) {
			if symbolForRepeat == 0 {
				return "", ErrInvalidString
			}
			count, e := strconv.Atoi(string(r))
			if e != nil {
				return "", ErrInvalidString
			}
			if isPreviousSymbolSlash {
				builder.WriteString(strings.Repeat(`\`+string(symbolForRepeat), count))
			} else {
				builder.WriteString(strings.Repeat(string(symbolForRepeat), count))
			}
			symbolForRepeat = 0
			isPreviousSymbolSlash = false
			continue
		}
		if r == '\\' {
			if isPreviousSymbolSlash {
				builder.WriteRune('\\')
			}
			isPreviousSymbolSlash = true
			continue
		}
		if !unicode.IsLetter(r) {
			return "", ErrInvalidString
		}
		if symbolForRepeat != 0 {
			builder.WriteRune(symbolForRepeat)
		}
		symbolForRepeat = r
	}
	if symbolForRepeat != 0 {
		builder.WriteRune(symbolForRepeat)
	}
	if isPreviousSymbolSlash {
		return "", ErrInvalidString
	}

	return builder.String(), nil
}
