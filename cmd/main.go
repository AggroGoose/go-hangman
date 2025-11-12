package main

import (
	"hangman/game"
)

func main() {
	gb := game.NewGame()

	gb.DisplayWord()

	gb = game.NewGame()

	gb.DisplayWord()
}