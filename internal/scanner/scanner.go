package scanner

import (
	"errors"
	"fmt"
	"regexp"
)

var (
	InvalidInputError  = errors.New("invalid input")
	InputTooLongError  = errors.New("input too long")
	InvalidLetterError = errors.New("invalid letter")
)

func ScanEngLetter() (rune, error) {
	for {
		line, err := ScanLine()
		if err != nil {
			return 0, err
		}

		if len(line) > 1 {
			return ' ', InputTooLongError
		}

		matched, err := regexp.MatchString("^[a-z]$", line)
		if err != nil || !matched {
			return rune(line[0]), InvalidLetterError
		}

		return rune(line[0]), nil
	}
}

func ScanLine() (line string, err error) {
	n, err := fmt.Scanln(&line)
	if err != nil || n != 1 {
		return line, InvalidInputError
	}
	return line, nil
}
