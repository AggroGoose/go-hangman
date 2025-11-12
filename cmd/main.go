package main

import (
	"fmt"
	"hangman/game"
)

func main() {
	isRunning := true

	for isRunning {
		fmt.Println("Welcome to Hangman!")
		fmt.Println("If you would like to start a new game type 'start' else enter 'quit': ")
		response := ""
		fmt.Scanln(&response)

		if response == "quit" {
			isRunning = false
		}
		if response == "start" {
			gb := game.NewGame()
			gb.StartGame()
		}
	}
}