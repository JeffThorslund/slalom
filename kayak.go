package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

type Racer struct {
	id   RacerId
	name string
}

type RacerId string

type EntryType int

const (
	Start EntryType = iota
	End
)

type Entry struct {
	racerId   RacerId
	time      time.Time
	entryType EntryType
}

/**
Assumptions: The data is in order of when it happened (validate for this).
*/

func main() {
	starts, ends, racers := parseCsvData("data/starts.csv", func(record []string) Entry {
		return Entry{
			RacerId(record[0]), parseTime(record[1]), Start,
		}
	}), parseCsvData("data/ends.csv", func(record []string) Entry {
		return Entry{
			RacerId(record[0]), parseTime(record[1]), End,
		}
	}), parseCsvData("data/racers.csv", func(record []string) Racer {
		return Racer{
			RacerId(record[0]), record[1],
		}
	})

	validationError := getValidationError(starts, ends, racers)

	if validationError != nil {
		log.Fatal(validationError)
	}

	// now we know (hopefully) the results are valid, we want to see the results in different views (think selectors)
	// we can continue without validation

	// 1. a "per person" breakdown of their races, sorted. This is the most natural way of constructing the structure so we start with that.
	sortedRacesPerRacer := createSortedEntriesPerRacer(starts, ends).ToRaces()

	// 2. a master list of sorted results, that can be filtered by catagory.
	var masterRaceList []Race
	for _, races := range sortedRacesPerRacer {
		masterRaceList = append(masterRaceList, races...)
	}

	sort.Slice(masterRaceList, func(i, j int) bool {
		return masterRaceList[i].getRaceTime() < masterRaceList[j].getRaceTime()
	})

	// 3. the "fun awards"

	fmt.Println(sortedRacesPerRacer, masterRaceList)
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
