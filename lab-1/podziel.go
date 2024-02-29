package main

import (
	"cmp"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	wordList := Read("slowa.txt")
	wordSet := map[string]bool{}
	for _, w := range wordList {
		wordSet[w] = true
	}
	wordCounter := map[string]int{}
	for _, w := range wordList {
		for _, parts := range Split(w) {
			AddIfIn(parts, wordSet, &wordCounter)
		}
	}
	pairs := Sort(wordCounter)
	for _, p := range pairs {
		fmt.Printf("%d %s\n", p.number, p.word)
	}
}

// Read wczytuje wyrazy z pliku o nazwie `filename` i zwraca je w
// wycinku tablicy lancuchow.
func Read(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return strings.Split(string(content), "\n")
}

// Split dzieli wyraz `word` na dwa niepuste lancuchy na wszystkie
// mozliwe sposoby i zwraca te dwa lancuchy w wycinku tablicy wycinkow
// tablic lancuchow.
func Split(word string) [][]string {
	r := [][]string{}
	for i := 1; i < len(word); i++ {
		r = append(r, []string{word[:i], word[i:]})
	}
	return r
}

// AddIfIn zwieksza licznik `counter` przy tych lancuchach z wycinka
// `parts`, ktore wystepuja w zbiorze lancuhow `set`.
func AddIfIn(parts []string, set map[string]bool, counter *map[string]int) {
	if set[parts[0]] && set[parts[1]] {
		(*counter)[parts[0]]++
		(*counter)[parts[1]]++
	}
}

type Pair struct {
	word   string
	number int
}

// Sort sortuje pary w malejacej kolejnosci pol `number`. Takie pary,
// ktore maja jednakowa wartosc pol `counter`, sortuje w kolejnosci
// leksykograficznej pol `word`

func Sort(counter map[string]int) []Pair {
	r := []Pair{}
	for w, c := range counter {
		r = append(r, Pair{w, c})
	}
	slices.SortFunc(r, func(a, b Pair) int {
		if n := b.number - a.number; n != 0 {
			return n
		}
		return cmp.Compare(a.word, b.word)
	})
	return r
}
