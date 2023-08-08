package main

import (
	"encoding/csv"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
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

func parseWords(f *os.File) (map[int]string, error) {
	reader := csv.NewReader(f)
	reader.Comma = '\t'
	reader.FieldsPerRecord = -1
	recs, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var p Pair
	pairs := make(map[int]string)
	for _, each := range recs {
		p.Roll, _ = strconv.Atoi(each[0])
		p.Word = each[1]
		pairs[p.Roll] = p.Word
	}
	return pairs, nil
}

func genRandomRoll(length int, min int, max int) int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	roll := 0
	for i := 0; i < length; i++ {
		digit := r.Intn(max-min+1) + min
		roll = roll*10 + digit
	}
	return roll
}

func getRandomWord(pairs map[int]string) string {
	// Get random int of length 4 with each digit from range [1, 6]
	return pairs[genRandomRoll(4, 1, 6)]
}

// getSecurePhrase generates a 'secure' passphrase consisting of 9 words
func getSecurePhrase(pairs map[int]string) string {
	phrase := ""
	for i := 0; i < 9; i++ {
		phrase += getRandomWord(pairs) + " "
	}
	return phrase[:len(phrase)-1]
}

// getPhrase generates a passphrase of length phraseLen. Note that this may not be long enough to be considered 'secure'.
func getPhrase(pairs map[int]string, phraseLen int) string {
	phrase := ""
	if phraseLen < 1 {
		return phrase
	}
	for i := 0; i < phraseLen; i++ {
		phrase += getRandomWord(pairs) + " "
	}
	return phrase[:len(phrase)-1]
}

func main() {
	words, err := os.Open("eff.org_files_2016_09_08_eff_short_wordlist_1.txt")
	check(err)
	defer words.Close()

	pairs, err := parseWords(words)
	check(err)

	fmt.Println(getSecurePhrase(pairs))
}
