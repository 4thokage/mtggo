package cmd

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"regexp"
)

const CsvCubeCobra = "CSV_CUBE_COBRA"
const CsvUrzaGatherer = "CSV_URZA_GATHERER"
const MTGO = "MTGO"

type MTGCard struct {
	Name    string
	Edition string
	Status  string
	Finish  string
	Number  string
	Count   string
}

func fromFile(deckFile string, fileType string) []MTGCard {
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

	if fileType == CsvCubeCobra {
		cards := fromCubeCobraCSV(f)
		return cards
	} else if fileType == CsvUrzaGatherer {
		cards := fromUrzaGathererCSV(f)
		return cards
	} else if fileType == MTGO {
		cards := fromMTGO(f)
		return cards
	} else {
		log.Fatal("No suitable reader for deck/collection file")
		return nil
	}

}

func fromCubeCobraCSV(f *os.File) []MTGCard {

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var cardList []MTGCard
	for i, line := range data {
		if i > 0 { // omit header line
			var rec MTGCard
			for j, field := range line {
				if j == 0 {
					rec.Name = field
				} else if j == 4 {
					rec.Edition = field
				} else if j == 5 {
					rec.Number = field
				} else if j == 8 {
					rec.Status = field
				} else if j == 9 {
					rec.Finish = field
				}
			}
			cardList = append(cardList, rec)
		}
	}

	return cardList
}

func fromUrzaGathererCSV(f *os.File) []MTGCard {

	csvReader := csv.NewReader(f)
	data, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	var cardList []MTGCard
	for i, line := range data {
		if i > 1 { // omit sep and header line
			var rec MTGCard
			for j, field := range line {
				if j == 0 {
					rec.Name = field
				} else if j == 9 {
					rec.Count = field
				}
			}
			cardList = append(cardList, rec)
		}
	}

	return cardList
}

func fromMTGO(f *os.File) []MTGCard {

	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	var lines []string
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}

	var cardList []MTGCard
	countRegex, err := regexp.Compile("^[^\\d]*(\\d+)")
	nameRegex, err := regexp.Compile("\\b([a-zA-Z,]+.*)")
	if err != nil {
		log.Fatal(err)
	}
	for _, text := range lines {
		var rec MTGCard
		rec.Count = countRegex.FindString(text)
		rec.Name = nameRegex.FindString(text)
		if rec.Name != "Sideboard" && rec.Name != "" {
			cardList = append(cardList, rec)
		}
	}

	return cardList
}
