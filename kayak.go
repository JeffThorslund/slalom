package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
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

	var entryParser = func(record []string) Entry {
		return Entry{
			Id(record[0]), parseTime(record[1]),
		}
	}

	starts := parseCsvData("data/starts.csv", entryParser)
	ends := parseCsvData("data/ends.csv", entryParser)

	validationError := validationErrorAggregator(
		assertValidRacerInformation(racers),
		assertNoDuplicateRacers(racers),
		assertOrderedRaceStarts(getTimes(starts)),
		assertEqualAmountOfStartsAndEnds(len(starts), len(ends)),
	)

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

func validationErrorAggregator(validationErrors ...error) error {
	for _, validationError := range validationErrors {
		if validationError != nil {
			return validationError
		}
	}

	return nil
}

func assertValidRacerInformation(racers []Racer) error {
	for i, racer := range racers {
		if racer.id == "" {
			return errors.New("Empty racer id." + "i:" + strconv.Itoa(i))
		}

		if racer.name == "" {
			return errors.New("Empty racer name." + "i:" + strconv.Itoa(i))
		}
	}

	return nil
}

func getTimes(starts []Entry) []time.Time {
	var times []time.Time

	for _, start := range starts {
		times = append(times, start.time)
	}

	return times
}

func assertOrderedRaceStarts(startTimes []time.Time) error {
	for i := 1; i < len(startTimes); i++ {
		if startTimes[i].Before(startTimes[i-1]) {
			return errors.New("Unordered start time at i:" + strconv.Itoa(i))
		}
	}
	return nil
}

func assertEqualAmountOfStartsAndEnds(starts int, ends int) error {

	if starts != ends {
		return errors.New("Number of starts:" + strconv.Itoa(starts) + "Number of ends:" + strconv.Itoa(ends))
	}

	return nil
}

func assertNoDuplicateRacers(racers []Racer) error {
	seenIds := make(map[Id]bool)
	seenNames := make(map[string]bool)

	for _, racer := range racers {
		if seenIds[racer.id] {
			return errors.New("Duplicate id:" + string(racer.id))
		}

		if seenNames[racer.name] {
			return errors.New("Duplicate name:" + racer.name)
		}
	}
	return nil
}
