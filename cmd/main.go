package main

import (
	"fmt"
	wb "hangman/data"
	"hangman/game"
)

func main() {
	word := wb.GetNewWord()

	fmt.Println(word)

	word = wb.GetNewWord()

	fmt.Println(word)

	game.HangStart()
	game.HangOne()
	game.HangTwo()
	game.HangThree()
	game.HangFour()
	game.HangFive()
}