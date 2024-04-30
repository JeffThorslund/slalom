package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"time"
)

type Racer struct {
	name string
	id   Id
}

type Id string

type Entry struct {
	id   Id
	time time.Time
}

type Race [2]time.Time // a start and end time

type Results struct {
}

/**
Assumptions: The data is in order of when it happened (validate for this).
All starts have ends. There are no ends that do not have starts.
*/

func main() {
	racers := parseCsvData("data/racers.csv", func(record []string) Racer {
		return Racer{
			id:   Id(record[0]),
			name: record[1],
		}
	})

	starts := parseCsvData("data/starts.csv", entryParser)
	ends := parseCsvData("data/ends.csv", entryParser)

	fmt.Println(starts, ends, racers)
}

func entryParser(record []string) Entry {
	return Entry{
		Id(record[0]), parseTime(record[1]),
	}
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

func validation() {

}
