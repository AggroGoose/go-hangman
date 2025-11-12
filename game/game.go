package game

import (
	"fmt"
	wb "hangman/data"
	"strings"
)
type GameBrain struct {
	word string
	score int
	remainingGuesses int
	isPlaying bool
	guesses []string
	currentState []string
}

func NewGame() GameBrain {
	word := wb.GetNewWord()
	var guesses []string
	var startState []string
	for i := 0; i < len(word); i++ {
		startState = append(startState, "_")
	}

	return GameBrain{
		word: word,
		score: 0,
		remainingGuesses: 5,
		isPlaying: false,
		guesses: guesses,
		currentState: startState,
	}
}

func (g GameBrain) DisplayWord(){
	fmt.Println(g.word)
	fmt.Println("Current word:", strings.Join(g.currentState, ""))
}
