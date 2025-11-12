package game

type GameBrain struct {
	word string
	score int
	isPlaying bool
	guesses []string
}