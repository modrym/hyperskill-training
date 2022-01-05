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
	var word string

	for {
		fmt.Scan(&word)

		if word == "exit" {
			fmt.Println("Bye!")
			break
		}

		fmt.Println(td.censorWord(word))
	}
}

func main() {
	tabooWords := TabooDB{}
	tabooWords.readPathFromInput()
	tabooWords.inputLoop()
}
