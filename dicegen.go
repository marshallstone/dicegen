package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Pair struct {
	Roll int
	Word string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	words, err := os.Open("eff.org_files_2016_09_08_eff_short_wordlist_1.txt")
	check(err)
	defer words.Close()

	reader := csv.NewReader(words)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1

	data, err := reader.ReadAll()
	check(err)

	var p Pair
	var pairs []Pair

	for _, each := range data {
		p.Roll, _ = strconv.Atoi(each[0])
		p.Word = each[1]
		pairs = append(pairs, p)
	}

	fmt.Print(pairs[1])
}
