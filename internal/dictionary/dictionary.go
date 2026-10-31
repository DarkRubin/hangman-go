package dictionary

import (
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
)

type Dictionary struct {
	words []string
	count int
}

func NewDictionary(filename string) (*Dictionary, error) {
	fileBytes, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("could not read '%s' file", filename)
	}

	words := strings.Fields(string(fileBytes))

	return &Dictionary{words, len(words)}, nil
}

func (dict Dictionary) RandomWord() string {
	return dict.words[rand.IntN(dict.count-1)]
}
