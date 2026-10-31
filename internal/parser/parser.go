package parser

import (
	"errors"
	"strings"
)

func ParseDecision(input string) (bool, error) {
	switch len(input) {
	case 1:
		switch strings.ToLower(input)[0] {
		case 'y':
			return true, nil
		case 'n':
			return false, nil
		}
	case 2:
		if strings.EqualFold(input, "no") {
			return false, nil
		}
	case 3:
		if strings.EqualFold(input, "yes") {
			return true, nil
		}
	}
	return false, errors.New("invalid input")
}
