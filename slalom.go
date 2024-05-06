package main

import (
	"fmt"
	"log"
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

	// parse the data from csv into arrays
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

	// validate the data without mutation, throw errors for humans to fix
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
