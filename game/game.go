package game

import (
	"fmt"
	wb "hangman/data"
	"regexp"
	"slices"
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

func (g *GameBrain) StartGame(){
	(*g).isPlaying = true
	g.NextGuess()
}

func (g *GameBrain) NextGuess(){
	r := regexp.MustCompile(`[a-zA-Z]`)

	for g.isPlaying {
		g.DisplayHang()
		fmt.Println("Word:", strings.Join(g.currentState, ""))
		fmt.Println("Lives Remaining:", g.remainingGuesses)
		fmt.Println("Currently Guessed:", strings.Join(g.guesses, ","))
		fmt.Println("Guess a letter a-z or enter 'quit' to stop:")
		currentGuess := ""
		fmt.Scanln(&currentGuess)
		if len(currentGuess) == 1 && r.MatchString(currentGuess) {
			match := g.evaluateGuess(strings.ToLower(currentGuess))
			if !match {
				fmt.Println("Whomp, whomp!")
				(*g).remainingGuesses--
			}
		}
		if currentGuess == "quit" {
			(*g).isPlaying = false
		}
		g.checkWinLoss()
	}
	fmt.Println("The Game is No Longer Playing")
}

func (g *GameBrain) evaluateGuess(guess string) bool {
	if slices.Contains(g.guesses, guess) {
			return true
		}
	(*g).guesses = append(g.guesses, guess)
	if !strings.Contains(g.word, guess) {
		return false
	}
	for i := 0; i < len(g.word); i++ {
		wordArr := []byte(g.word)
		if wordArr[i] == []byte(guess)[0] {
			(*g).currentState[i] = guess
		}
	}
	fmt.Println("You guessed one!")
	return true
}

func (g *GameBrain) checkWinLoss() {
	if g.remainingGuesses < 1 {
		g.loseGame()
		(*g).isPlaying = false
	}
	if g.word == strings.Join(g.currentState, "") {
		g.winGame()
		(*g).isPlaying = false
	}
}

func (g GameBrain) winGame() {
	g.DisplayHang()
	fmt.Println("Word:", strings.Join(g.currentState, ""))
	fmt.Println("Lives Remaining:", g.remainingGuesses)
	fmt.Println("Currently Guessed:", strings.Join(g.guesses, ","))
	fmt.Println("CONGRATULATIONS YOU WON!!!")
}

func (g GameBrain) loseGame() {
	HangFive()
	fmt.Println("Word:", g.word)
	fmt.Println("Lives Remaining:", g.remainingGuesses)
	fmt.Println("Currently Guessed:", strings.Join(g.guesses, ","))
	fmt.Println("SORRY, YOU LOST! PLEASE TRY AGAIN!")
}

func (g GameBrain) DisplayHang() {
	if g.remainingGuesses == 5 {
		HangStart()
	}
	if g.remainingGuesses == 4 {
		HangOne()
	}
	if g.remainingGuesses == 3 {
		HangTwo()
	}
	if g.remainingGuesses == 2 {
		HangThree()
	}
	if g.remainingGuesses == 1 {
		HangFour()
	}
}