package parsing

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/JeffThorslund/slalom-results/entry"
	"github.com/JeffThorslund/slalom-results/racer"
)

func ProcessRawData() ([]entry.Entry, []entry.Entry, []racer.Racer) {
	starts := parseCsvData("testdata/starts.csv", parseStart)
	ends := parseCsvData("testdata/ends.csv", parseEnd)
	racers := parseCsvData("testdata/racers.csv", parseRacer)

	return starts, ends, racers
}

func parseCsvData[D interface{}](path string, parser func([]string) D) []D {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal("Error while reading file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}

	var results []D

	for _, record := range records[1:] {
		results = append(results, parser(record))
	}

	return results
}

func parseStart(record []string) entry.Entry {
	return entry.NewEntry(
		racer.RacerId(record[0]),
		parseTime(record[1]),
		entry.Start,
	)
}

func parseEnd(record []string) entry.Entry {
	return entry.NewEntry(
		racer.RacerId(record[0]),
		parseTime(record[1]),
		entry.End,
	)
}

func parseTime(timeStr string) time.Time {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func parseRacer(record []string) racer.Racer {
	return racer.NewRacer(
		racer.RacerId(record[0]),
		record[1],
		parseGender(record[2]),
		parseCategory(record[3]),
	)
}

func parseCategory(s string) racer.Category {
	var category racer.Category
	var err error

	switch s {
	case "b":
		category, err = racer.Beginner, nil
	case "i":
		category, err = racer.Intermediate, nil
	case "a":
		category, err = racer.Advanced, nil
	default:
		category, err = 0, fmt.Errorf("invalid catergory: %s", s)
	}

	if err != nil {
		log.Fatal(err)
	}

	return category
}

func parseGender(s string) racer.Gender {
	var gender racer.Gender
	var err error

	switch s {
	case "m":
		gender, err = racer.Male, nil
	case "f":
		gender, err = racer.Female, nil
	default:
		gender, err = 0, fmt.Errorf("invalid gender: %s", s)
	}

	if err != nil {
		log.Fatal(err)
	}

	return gender
}
