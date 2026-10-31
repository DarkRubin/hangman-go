package game

import (
	"slices"
	"strings"
)

const MaxMistakesCount = 6

type GameState struct {
	word           string
	mask           []rune
	guessedLetters []rune
	correctGuesses int
	mistakes       int
}

func New(word string) *GameState {

	mask := strings.Repeat("*", len(word))
	return &GameState{
		word:           word,
		mask:           []rune(mask),
		guessedLetters: []rune{},
	}
}

func (state *GameState) IsAlreadyGuessed(letter rune) bool {
	return slices.Contains(state.guessedLetters, letter)
}

func (state *GameState) Guess(letter rune) (success bool) {
	state.guessedLetters = append(state.guessedLetters, letter)
	if strings.ContainsRune(state.word, letter) {
		for i, symbol := range []rune(state.word) {
			if symbol == letter {
				state.mask[i] = symbol
				state.correctGuesses = state.correctGuesses + 1
			}
		}
		return true
	}

	state.mistakes = state.mistakes + 1
	return false
}

func (state *GameState) IsGameEnded() bool {
	return state.mistakes == MaxMistakesCount || state.IsWin()
}

func (state *GameState) IsWin() bool {
	return state.correctGuesses == len(state.word)
}

func (state *GameState) GetMask() string {
	formatedMask := strings.Builder{}
	for _, letter := range state.mask {
		formatedMask.WriteRune(letter)
		formatedMask.WriteRune(' ')
	}
	return formatedMask.String()
}

func (state *GameState) GetMistakes() int {
	return state.mistakes
}
