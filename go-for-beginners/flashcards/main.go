package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Question struct {
	term, definition string
}

func readValue(reader *bufio.Reader) string {
	line, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("there was a problem when reading from stdin")
	}

	return strings.TrimSpace(line)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Input the number of cards:")
	number, err := strconv.Atoi(readValue(reader))

	if err != nil {
		log.Fatal("the number is incorrect")
	}

	db := make([]Question, 0, number)

	// DB creation stage
	for i := 1; i <= number; i++ {
		fmt.Printf("The term for card #%d:\n", i)
		term := readValue(reader)

		fmt.Printf("The definition for card #%d:\n", i)
		definition := readValue(reader)

		db = append(db, Question{term, definition})
	}

	// test stage
	for _, val := range db {
		fmt.Printf("Print the definition of \"%s\":\n", val.term)

		if readValue(reader) == val.definition {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("Wrong. The right answer is \"%s\".\n", val.definition)
		}
	}
}
