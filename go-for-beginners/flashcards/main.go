package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Qdb struct {
	reader             *bufio.Reader
	dbTermToDefinition map[string]string
}

func (q *Qdb) readValue() string {
	line, err := q.reader.ReadString('\n')

	if err != nil {
		log.Fatal("there was a problem when reading from stdin")
	}

	return strings.TrimSpace(line)
}

func NewQdb() *Qdb {
	q := Qdb{}
	q.reader = bufio.NewReader(os.Stdin)
	q.dbTermToDefinition = make(map[string]string)

	return &q
}

func (q *Qdb) getUniqueValue(valueName string, checkDuplicateFunc func(string) bool) string {
	var value string

	for {
		value = q.readValue()

		if checkDuplicateFunc(value) {
			fmt.Printf("The %s \"%s\" already exists. Try again:\n", valueName, value)
			continue
		}

		break
	}

	return value
}

func (q *Qdb) dbAddQuestion(term, definition string) {
	q.dbTermToDefinition[term] = definition
}

func (q *Qdb) dbRemoveQuestion(term string) bool {
	if _, ok := q.dbTermToDefinition[term]; !ok {
		return false
	}

	delete(q.dbTermToDefinition, term)

	return true
}

func (q *Qdb) actionAdd() {
	fmt.Println("The card:")

	term := q.getUniqueValue("card", func(s string) bool {
		_, ok := q.dbTermToDefinition[s]
		return ok
	})

	fmt.Println("The definition of the card:")

	definition := q.getUniqueValue("definition", func(s string) bool {
		for _, value := range q.dbTermToDefinition {
			if s == value {
				return true
			}
		}
		return false
	})

	q.dbAddQuestion(term, definition)

	fmt.Printf("The pair (\"%s\":\"%s\") has been added.\n", term, definition)
}

func (q *Qdb) actionRemove() {
	fmt.Println("Which card?")

	term := q.readValue()
	if !q.dbRemoveQuestion(term) {
		fmt.Println("Can't remove \"%s\": there is no such card.", term)
		return
	}

	fmt.Println("The card has been removed.")

}

func (q *Qdb) actionImport() {
	fmt.Println("File name:")

	fname := q.readValue()

	file, err := os.Open(fname)

	if err != nil {
		if os.IsExist(err) {
			fmt.Println("File not found.")
			return
		}
		fmt.Printf("Error opening the file: %s\n", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNo := 0
	twoLines := make([]string, 2, 2)

	for scanner.Scan() {
		twoLines[lineNo%2] = strings.TrimSpace(scanner.Text())

		if lineNo%2 == 1 {
			q.dbAddQuestion(twoLines[0], twoLines[1])
		}

		lineNo++
	}

	fmt.Printf("%d cards have been loaded.\n", (lineNo+1)/2)
}

func (q *Qdb) actionExport() {
	fmt.Println("File name:")

	fname := q.readValue()

	file, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	if err != nil {
		fmt.Println("Could not save the file.")
		return
	}

	writer := bufio.NewWriter(file)

	for key, value := range q.dbTermToDefinition {
		writer.WriteString(key + "\n" + value + "\n")
	}

	if writer.Flush() != nil || file.Close() != nil {
		fmt.Println("Could not save the file.")
		return
	}

	fmt.Printf("%d cards have been saved.", len(q.dbTermToDefinition))
}

func (q *Qdb) actionAsk() {
	fmt.Println("How many times to ask?")
	number, err := strconv.Atoi(q.readValue())

	if err != nil {
		fmt.Println("Wrong number.")
		return
	}

	for key, value := range q.dbTermToDefinition {
		if number <= 0 {
			break
		}

		fmt.Printf("Print the definition of \"%s\":\n", key)

		answer := q.readValue()

		if answer == value {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("Wrong. The right answer is \"%s\"", value)

			for _, otherValue := range q.dbTermToDefinition {
				if otherValue == answer {
					fmt.Printf(", but your definition is correct for \"%s\"", otherValue)
					break
				}
			}
			fmt.Println(".")
		}

		number--
	}
}

func (q *Qdb) ActionLoop() {
	for {
		fmt.Println("Input the action (add, remove, import, export, ask, exit):")

		switch q.readValue() {
		case "add":
			q.actionAdd()
		case "remove":
			q.actionRemove()
		case "import":
			q.actionImport()
		case "export":
			q.actionExport()
		case "exit":
			fmt.Println("Bye bye!")
			return
		default:
			fmt.Println("Wrong command!")
		}

		fmt.Println()
	}
}

func main() {
	qdb := NewQdb()
	qdb.ActionLoop()
}
