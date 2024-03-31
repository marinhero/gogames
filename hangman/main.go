package main

import (
	"fmt"
	"math/rand"
)

// Game Plan
// Hangman is a guessing game. You show a word, and then accept and input of a character.
// If the character exists within the word, show the character, if not, mark it as a strike.
// Word: Hangman
// Game Example:
// H___m_n
// > a
// Ha__man
// > k

// Print Game State
//   - Print word you are guessing
//   - Print hangman state
// Derive word?
// Read user input
// - validate (only letters, no words, no numbers)
// Determine outcome and update letters or hangman state.

const StrikeLimit = 10

var WordList = [8]string{
	"ASERE",
	"ZOMBIE",
	"GOPHER",
	"GODZILLA",
	"GOAT",
	"MEXICO",
	"PROGRAMMING",
	"TZINTZUNTZAN",
}

func getRandomWord() string {
	randomNumber := rand.Intn(7)
	return WordList[randomNumber]
}

func welcome() {
	fmt.Println("Welcome to Hangman!")
	fmt.Println(fmt.Sprintf("You have %d attemps to guess the following word:", StrikeLimit))
}

func printGameState(targetWord string, guessedLetters map[rune]bool) {
	for i := 0; i < len(targetWord); i++ {
		if guessedLetters[rune(targetWord[i])] {
			fmt.Print(string(targetWord[i]))
		} else {
			fmt.Print("_")
		}
	}
}

func main() {
	welcome()
	targetWord := getRandomWord()
	guessedLetters := map[rune]bool{}

	guessedLetters[rune(targetWord[0])] = true
	guessedLetters[rune(targetWord[len(targetWord)-1])] = true

	printGameState(targetWord, guessedLetters)

}
