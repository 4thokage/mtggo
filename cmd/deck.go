package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func cardNamesFromFile(deckFile string) {
	f, err := os.Open(deckFile)
	if err != nil {
		log.Fatal(err)
	}
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		exp := regexp.MustCompile(`^[0-9]+`)

		matches := exp.FindAllString(scanner.Text(), -1)
		fmt.Printf("line: %s\n", scanner.Text())
		log.Println(matches)

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
