package main

import (
	"fmt"
	wb "hangman/data"
)

func main() {
	word := wb.GetNewWord()

	fmt.Println(word)

	word = wb.GetNewWord()

	fmt.Println(word)
}