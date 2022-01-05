package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const CensorString = "*"

type TabooDB struct {
	db map[string]struct{}
}

func (td *TabooDB) readPathFromInput() {
	var path string
	fmt.Scan(&path)

	file, err := os.Open(path)

	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(file)

	td.db = make(map[string]struct{})

	for scanner.Scan() {
		td.db[strings.ToLower(scanner.Text())] = struct{}{}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (td *TabooDB) isTaboo(word string) bool {
	_, ok := td.db[strings.ToLower(word)]

	return ok
}

func (td *TabooDB) censorWord(word string) string {
	if !td.isTaboo(word) {
		return word
	}

	return strings.Repeat(CensorString, len(word))
}

func (td *TabooDB) inputLoop() {
	var word, dot string
	var wordsSplit []string

	for {
		fmt.Scanln(&word)

		if word == "exit" {
			fmt.Println("Bye!")
			break
		}

		if word[len(word)-1] == '.' {
			dot = "."
		} else {
			dot = ""
		}

		word = strings.TrimRight(word, ".")
		wordsSplit = strings.Split(word, " ")

		for index, val := range wordsSplit {
			wordsSplit[index] = td.censorWord(val)
		}

		fmt.Println(strings.Join(wordsSplit, " ") + dot)
	}
}

func main() {
	tabooWords := TabooDB{}
	tabooWords.readPathFromInput()
	tabooWords.inputLoop()
}
