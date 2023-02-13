package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const maxTurns = 6

func main() {
	// define the word to be guessed
	word := "hello"

	// initialize a slice to store the guessed letters
	lettersGuessed := make([]string, 0)

	// initialize the number of turns taken
	turns := 0

	// create a flag to indicate if the word has been fully guessed
	wordGuessed := false

	// loop until the word is fully guessed or the maximum number of turns is reached
	for turns < maxTurns && !wordGuessed {
		// display the word with underscores for unguessed letters
		display := ""
		for _, char := range word {
			if strings.Contains(strings.Join(lettersGuessed, ""), string(char)) {
				display += string(char)
			} else {
				display += "_"
			}
		}

		// display the current word and the number of turns taken
		fmt.Println("Current word:", display)
		fmt.Println("Turns taken:", turns)

		// prompt the user to guess a letter
		fmt.Print("Guess a letter: ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// add the letter to the list of guessed letters
		lettersGuessed = append(lettersGuessed, input)

		// check if the letter is in the word
		if !strings.Contains(word, input) {
			turns++
		}

		// check if the word has been fully guessed
		if strings.Join(lettersGuessed, "") == word {
			wordGuessed = true
		}
	}

	// display a message indicating if the word was fully guessed or if the maximum number of turns was reached
	if wordGuessed {
		fmt.Println("Congratulations! You fully guessed the word.")
	} else {
		fmt.Println("Sorry, you ran out of turns. The word was", word)
	}
}
