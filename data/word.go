package data

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
)

func GetNewWord() string {
	file, err := os.Open("../data/wordbank.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)

	records, _ := reader.ReadAll()

	word := records[rand.Intn(len(records))][0]

	if len(word) > 0 {
		return word
	} else {
		return "ganso"
	}
}