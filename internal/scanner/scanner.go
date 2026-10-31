package scanner

import (
	"errors"
	"fmt"
	"io"
	"regexp"
)

var (
	ErrEOF           = errors.New("EOF")
	ErrInvalidInput  = errors.New("invalid input")
	ErrInputTooLong  = errors.New("input too long")
	ErrInvalidLetter = errors.New("invalid letter")
)

func ScanEngLetter() (rune, error) {
	for {
		line, err := ScanLine()
		if err != nil {
			return 0, err
		}

		if len(line) > 1 {
			return ' ', ErrInputTooLong
		}

		matched, err := regexp.MatchString("^[a-z]$", line)
		if err != nil || !matched {
			return rune(line[0]), ErrInvalidLetter
		}

		return rune(line[0]), nil
	}
}

func ScanLine() (line string, err error) {
	n, err := fmt.Scanln(&line)
	if err != nil || n != 1 {
		if err == io.EOF {
			return "", ErrEOF
		}
		return line, ErrInvalidInput
	}
	return line, nil
}
