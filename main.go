package main

import (
	"errors"
	"fmt"
	"hangman/internal/dictionary"
	"hangman/internal/game"
	"hangman/internal/parser"
	"hangman/internal/scanner"
)

const WordStorageFilename = "embed/words.txt"

var dict *dictionary.Dictionary

func main() {
	fmt.Println("Welcome to Hangman!")
	fmt.Println("Do you want to play? 'Yes' or 'No': ")

	for {
		input, err := scanner.ScanLine()
		if err != nil {
			fmt.Println("Empty or missing input. Enter 'Y' or 'N'")
			continue
		}

		userWantsPlay, err := parser.ParseDecision(input)
		if err != nil {
			fmt.Println("Unknown input. Enter 'Y' or 'N'")
			continue
		}

		if userWantsPlay {
			if err = play(); err != nil {
				fmt.Println("Error. Could not start game")
				return
			}
			fmt.Println("Do you want to play again? 'Y' / 'n'")
		} else {
			fmt.Println("Goodbye")
			return
		}
	}
}

func play() (err error) {
	if dict == nil {
		dict, err = dictionary.NewDictionary(WordStorageFilename)
		if err != nil {
			return err
		}
	}

	state := game.New(dict.RandomWord())

	for !state.IsGameEnded() {
		fmt.Println("Guess the word: ", state.GetMask())
		fmt.Println(state.GallowsArt())
		fmt.Println("Current mistake count: ", state.GetMistakes())

		letter := nextUniqueLetter(state)
		if !state.Guess(letter) {
			fmt.Printf("Unluck. There no '%c', in word '%s'! Mistakes count increased\n",
				letter, state.GetMask())
		}
	}

	fmt.Println(state.GallowsArt())
	if state.IsWin() {
		fmt.Println("You won! Congrats!")
	} else {
		fmt.Println("You lost! Congrats!")
	}

	return nil
}

func nextUniqueLetter(state *game.GameState) rune {
	for {
		letter, err := scanner.ScanEngLetter()
		if err == nil {
			if state.IsAlreadyGuessed(letter) {
				fmt.Printf("Letter '%c' was already guessed, try another one\n", letter)
				continue
			}
			return letter
		}

		switch {
		case errors.Is(err, scanner.InvalidInputError):
			fmt.Println("Empty or missing input. Enter enter one letter from English alphabet")
		case errors.Is(err, scanner.InputTooLongError):
			fmt.Println("You can choose only one English letter in one try")
		case errors.Is(err, scanner.InvalidLetterError):
			fmt.Printf("Incorrect letter '%c'. Enter enter one letter from English alphabet\n", letter)
		default:
			panic("unreachable")
		}
	}
}
