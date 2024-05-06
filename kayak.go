package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type Racer struct {
	id   Id
	name string
}

type Id string

type EntryType int

const (
	Start EntryType = iota
	End
)

type Entry struct {
	racerId   Id
	time      time.Time
	entryType EntryType
}

/**
Assumptions: The data is in order of when it happened (validate for this).
*/

func main() {

	starts := parseCsvData("data/starts.csv", func(record []string) Entry {
		return Entry{
			Id(record[0]), parseTime(record[1]), Start,
		}
	})

	ends := parseCsvData("data/ends.csv", func(record []string) Entry {
		return Entry{
			Id(record[0]), parseTime(record[1]), End,
		}
	})

	racers := parseCsvData("data/racers.csv", func(record []string) Racer {
		return Racer{
			Id(record[0]), record[1],
		}
	})

	validationError := getValidationError(starts, ends, racers)

	if validationError != nil {
		log.Fatal(validationError)
	}

	fmt.Println(starts, ends, racers)
}

func parseTime(timeStr string) time.Time {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		log.Fatal(err)
	}
	return t
}

func parseCsvData[D Racer | Entry](path string, parser func([]string) D) []D {
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
