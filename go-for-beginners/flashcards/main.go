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
	dbDefinitionToTerm map[string]string
	dbOrder            []*string
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
	q.dbDefinitionToTerm = make(map[string]string)

	return &q
}

func (q *Qdb) FillDb() {
	fmt.Println("Input the number of cards:")
	number, err := strconv.Atoi(q.readValue())

	if err != nil {
		log.Fatal("the number is incorrect")
	}

	for i := 1; i <= number; i++ {
		q.AddQuestion(i)
	}
}

func (q *Qdb) getUniqueValue(valueName string, dataMap *map[string]string) string {
	var value string

	for {
		value = q.readValue()

		if _, ok := (*dataMap)[value]; ok {
			fmt.Printf("The %s \"%s\" already exists. Try again:\n", valueName, value)
			continue
		}

		break
	}

	return value
}

func (q *Qdb) AddQuestion(number int) {
	fmt.Printf("The term for card #%d:\n", number)

	term := q.getUniqueValue("term", &q.dbTermToDefinition)

	fmt.Printf("The definition for card #%d:\n", number)

	definition := q.getUniqueValue("definition", &q.dbDefinitionToTerm)

	q.dbTermToDefinition[term] = definition
	q.dbDefinitionToTerm[definition] = term
	q.dbOrder = append(q.dbOrder, &term)
}

func (q *Qdb) Test() {
	for _, termPtr := range q.dbOrder {
		term := *termPtr
		definition := q.dbTermToDefinition[term]

		fmt.Printf("Print the definition of \"%s\":\n", term)

		userAnswer := q.readValue()

		if userAnswer == definition {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("Wrong. The right answer is \"%s\"", definition)

			if value, ok := q.dbDefinitionToTerm[userAnswer]; ok {
				fmt.Printf(", but your definition is correct for \"%s\"", value)
			}
			fmt.Println(".")
		}
	}
}

func main() {
	qdb := NewQdb()
	qdb.FillDb()
	qdb.Test()
}
