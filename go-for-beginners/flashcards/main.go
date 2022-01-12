package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type any = interface{}

type customLogger struct {
	builder strings.Builder
}

var Log customLogger

func (l *customLogger) Print(s string) {
	l.builder.WriteString(s)
	fmt.Print(s)
}

func (l *customLogger) Println(s string) {
	l.builder.WriteString(s + "\n")
	fmt.Println(s)
}

func (l *customLogger) Printf(s string, args ...any) {
	s = fmt.Sprintf(s, args...)
	l.builder.WriteString(s)
	fmt.Print(s)
}

func (l *customLogger) Printfln(s string, args ...any) {
	s = fmt.Sprintf(s+"\n", args...)
	l.builder.WriteString(s)
	fmt.Print(s)
}

func (l *customLogger) SaveFile(name string) error {
	file, err := os.OpenFile(name, os.O_WRONLY|os.O_CREATE, 0644)

	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(l.builder.String())

	err = writer.Flush()

	if err != nil {
		return err
	}

	return file.Close()
}

type Def struct {
	word     string
	mistakes int
}

type Qdb struct {
	reader             *bufio.Reader
	dbTermToDefinition map[string]Def
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
	q.dbTermToDefinition = make(map[string]Def)

	return &q
}

func (q *Qdb) getUniqueValue(valueName string, checkDuplicateFunc func(string) bool) string {
	var value string

	for {
		value = q.readValue()

		if checkDuplicateFunc(value) {
			Log.Printf("The %s \"%s\" already exists. Try again:\n", valueName, value)
			continue
		}

		break
	}

	return value
}

func (q *Qdb) dbAddQuestion(term, definition string, mistakes int) {
	q.dbTermToDefinition[term] = Def{definition, mistakes}
}

func (q *Qdb) dbRemoveQuestion(term string) bool {
	if _, ok := q.dbTermToDefinition[term]; !ok {
		return false
	}

	delete(q.dbTermToDefinition, term)

	return true
}

func (q *Qdb) actionAdd() {
	Log.Println("The card:")

	term := q.getUniqueValue("card", func(s string) bool {
		_, ok := q.dbTermToDefinition[s]
		return ok
	})

	Log.Println("The definition of the card:")

	definition := q.getUniqueValue("definition", func(s string) bool {
		for _, value := range q.dbTermToDefinition {
			if s == value.word {
				return true
			}
		}
		return false
	})

	q.dbAddQuestion(term, definition, 0)

	Log.Printf("The pair (\"%s\":\"%s\") has been added.\n", term, definition)
}

func (q *Qdb) actionRemove() {
	Log.Println("Which card?")

	term := q.readValue()
	if !q.dbRemoveQuestion(term) {
		Log.Printf("Can't remove \"%s\": there is no such card.\n", term)
		return
	}

	Log.Println("The card has been removed.")

}

func (q *Qdb) actionImport() {
	Log.Println("File name:")

	fname := q.readValue()

	file, err := os.Open(fname)

	if err != nil {
		Log.Println("File not found.")
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	lineNo := 0
	threeLines := make([]string, 3, 3)

	for scanner.Scan() {
		threeLines[lineNo%3] = strings.TrimSpace(scanner.Text())

		if lineNo%3 == 1 {
			mistakes, err := strconv.Atoi(threeLines[2])

			if err != nil {
				mistakes = 0
			}

			q.dbAddQuestion(threeLines[0], threeLines[1], mistakes)
		}

		lineNo++
	}

	Log.Printf("%d cards have been loaded.\n", (lineNo+1)/3)
}

func (q *Qdb) actionExport() {
	Log.Println("File name:")

	fname := q.readValue()

	file, err := os.OpenFile(fname, os.O_CREATE|os.O_WRONLY, 0644)
	defer file.Close()

	if err != nil {
		Log.Println("Could not save the file.")
		return
	}

	writer := bufio.NewWriter(file)

	for key, value := range q.dbTermToDefinition {
		writer.WriteString(key + "\n" + value.word + "\n" + strconv.Itoa(value.mistakes) + "\n")
	}

	if writer.Flush() != nil || file.Close() != nil {
		Log.Println("Could not save the file.")
		return
	}

	Log.Printf("%d cards have been saved.", len(q.dbTermToDefinition))
}

func (q *Qdb) actionAsk() {
	Log.Println("How many times to ask?")
	number, err := strconv.Atoi(q.readValue())

	if err != nil {
		Log.Println("Wrong number.")
		return
	}

	for number > 0 {
		for key, value := range q.dbTermToDefinition {
			if number <= 0 {
				break
			}

			Log.Printf("Print the definition of \"%s\":\n", key)

			answer := q.readValue()

			if answer == value.word {
				Log.Println("Correct!")
			} else {
				Log.Printf("Wrong. The right answer is \"%s\"", value.word)
				value.mistakes++
				q.dbTermToDefinition[key] = value

				for _, otherValue := range q.dbTermToDefinition {
					if otherValue.word == answer {
						Log.Printf(", but your definition is correct for \"%s\"", otherValue.word)
						break
					}
				}
				Log.Println(".")
			}

			number--
		}
	}
}

func (q *Qdb) actionLog() {
	Log.Println("File name:")

	fileName := q.readValue()

	err := Log.SaveFile(fileName)
	if err == nil {
		Log.Println("The log has been saved.")
	} else {
		Log.Printfln("Error: %s", err)
	}
}

func (q *Qdb) actionHardestCard() {
	var biggest struct {
		words    []string
		mistakes int
	}

	for key, value := range q.dbTermToDefinition {
		if len(biggest.words) == 0 {
			biggest.words = append(biggest.words, key)
			biggest.mistakes = value.mistakes
			continue
		}

		if biggest.mistakes < value.mistakes {
			biggest.mistakes = value.mistakes
			biggest.words = append([]string(nil), value.word)
		} else if biggest.mistakes == value.mistakes {
			biggest.words = append(biggest.words, value.word)
		}
	}

	if biggest.mistakes == 0 {
		Log.Println("There are no cards with errors.")
		return
	}

	Log.Print("The hardest card")

	if len(biggest.words) == 1 {
		Log.Print(" is \"")
	} else {
		Log.Print("s are \"")
	}

	Log.Printf("%s\". You have %d errors answering ",
		strings.Join(biggest.words, "\", \""), biggest.mistakes)

	if len(biggest.words) == 1 {
		Log.Println("it.")
	} else {
		Log.Println("them.")
	}

}

func (q *Qdb) actionResetStats() {
	for key, value := range q.dbTermToDefinition {
		value.mistakes = 0

		q.dbTermToDefinition[key] = value
	}

	Log.Println("Card statistics have been reset.")
}

func (q *Qdb) ActionLoop() {
	for {
		Log.Println("Input the action (add, remove, import, export, ask, " +
			"exit, log, hardest card, reset stats):")

		switch q.readValue() {
		case "add":
			q.actionAdd()
		case "remove":
			q.actionRemove()
		case "import":
			q.actionImport()
		case "export":
			q.actionExport()
		case "ask":
			q.actionAsk()
		case "log":
			q.actionLog()
		case "hardest card":
			q.actionHardestCard()
		case "reset stats":
			q.actionResetStats()
		case "exit":
			Log.Println("Bye bye!")
			return
		default:
			Log.Println("Wrong command!")
		}

		Log.Println("")
	}
}

func main() {
	qdb := NewQdb()
	qdb.ActionLoop()
}
