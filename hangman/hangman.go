package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"unicode"
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

const StrikeLimit = 9

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

func welcomeBanner() {
	fmt.Println("Welcome to Hangman!")
	fmt.Println(fmt.Sprintf("You have %d attemps to guess the following word:", StrikeLimit))
}

func printGameState(targetWord string, guessedLetters map[rune]bool) bool {
	endGame := false
	for i := 0; i < len(targetWord); i++ {
		currentChar := string(targetWord[i])
		if currentChar == " " {
			fmt.Println(" ")
		}

		if guessedLetters[unicode.ToLower(rune(targetWord[i]))] {
			fmt.Print(currentChar)
		} else {
			fmt.Print("_")
			endGame = true
		}
	}
	fmt.Println()
	return endGame
}

func printStrikes(state int) {
	filePath := fmt.Sprintf("./hangStates/%d.txt", state)
	data, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	fmt.Println(string(data))
}

func checkChar(s string, c rune) bool {
	var current rune
	for _, char := range s {
		current = unicode.ToLower(char)
		if current == c {
			return true
		}
	}
	return false
}

func isAlreadyRegistered(input rune, preChecked *map[rune]bool) bool {
	if (*preChecked)[input] {
		return true
	}
	return false
}

func gameLoop(targetWord string, guessedLetters *map[rune]bool) {
	errorCount := 0

	for errorCount < StrikeLimit {
		fmt.Print("> ")
		var input string
		_, err := fmt.Scanln(&input)
		if err != nil {
			return
		}

		parsedInput := unicode.ToLower(rune(input[0]))
		if !isAlreadyRegistered(parsedInput, guessedLetters) {
			if checkChar(targetWord, parsedInput) {
				(*guessedLetters)[parsedInput] = true
			} else {
				errorCount++
				printStrikes(errorCount)
			}
		}

		if !printGameState(targetWord, *guessedLetters) {
			fmt.Println("VICTORY")
			return
		}
	}

	if errorCount == StrikeLimit {
		fmt.Println("YOU ARE DEAD")
	}
}

func main() {
	welcomeBanner()
	targetWord := getRandomWord()
	guessedLetters := map[rune]bool{}

	firstChar := unicode.ToLower(rune(targetWord[0]))
	lastChar := unicode.ToLower(rune(targetWord[len(targetWord)-1]))

	guessedLetters[firstChar] = true
	guessedLetters[lastChar] = true

	printGameState(targetWord, guessedLetters)

	gameLoop(targetWord, &guessedLetters)
}
