package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
)

func main() {
	file, err := os.Open("wordbank.csv")
	if err != nil {
		fmt.Println(err)
	}
	reader := csv.NewReader(file)

	records, _ := reader.ReadAll()

	fmt.Println(records)

	word := records[rand.Intn(len(records))][0]

	fmt.Println(word)
}