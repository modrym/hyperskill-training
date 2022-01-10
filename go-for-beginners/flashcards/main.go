package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func readValue(reader *bufio.Reader) string {
	line, err := reader.ReadString('\n')

	if err != nil {
		log.Fatal("there was a problem when reading from stdin")
	}

	return strings.TrimSpace(line)
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	readValue(reader) // discard one value

	if readValue(reader) == readValue(reader) {
		fmt.Println("Your answer is right!")
	} else {
		fmt.Println("Your answer is wrong...")
	}
}
