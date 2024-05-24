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

func ProcessRawData() ([]entry.Entry, []entry.Entry, racer.Racers) {
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

	gender, err := parseGender(record[2])

	if err != nil {
		log.Fatal(err, record)
	}

	category, err := parseCategory(record[3])

	if err != nil {
		log.Fatal(err, record)
	}

	return racer.NewRacer(
		racer.RacerId(record[0]),
		record[1],
		gender,
		category,
	)
}

func parseCategory(s string) (racer.Category, error) {

	switch s {
	case "b":
		return racer.Beginner, nil
	case "i":
		return racer.Intermediate, nil
	case "a":
		return racer.Advanced, nil
	default:
		return 0, fmt.Errorf("invalid catergory: %s", s)
	}

}

func parseGender(s string) (racer.Gender, error) {
	switch s {
	case "m":
		return racer.Male, nil
	case "f":
		return racer.Female, nil
	default:
		return 0, fmt.Errorf("invalid gender: %s", s)
	}
}
