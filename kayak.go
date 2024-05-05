package main

import (
	"encoding/csv"
	"errors"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
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
All starts have ends. There are no ends that do not have starts.
*/

func main() {
	racers := parseCsvData("data/racers.csv", func(record []string) Racer {
		return Racer{
			Id(record[0]), record[1],
		}
	})

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

	validationError := validationErrorAggregator(
		assertValidRacerInformation(racers),
		assertNoDuplicateRacers(racers),
		assertOrderedRaceStarts(getTimes(starts)),
		assertThatAllRacesEnd(starts, ends),
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

func assertThatAllRacesEnd(starts []Entry, ends []Entry) error {

	allRaces := append(starts, ends...)

	sort.Slice(allRaces, func(i, j int) bool {
		return allRaces[i].time.Before(allRaces[j].time)
	})

	raceMap := make(map[Id][]Entry)

	for _, race := range allRaces {
		raceMap[race.racerId] = append(raceMap[race.racerId], race)
	}

	for _, entries := range raceMap {

		// track if a user is racing
		isRacing := false

		for _, entry := range entries {
			if !isRacing && (entry.entryType == Start) { // Racer starts a race
				isRacing = true
			} else if isRacing && (entry.entryType == End) { // Racer ends a race
				isRacing = false
			} else if !isRacing && (entry.entryType == End) { // Racer tries to end but is not currently racing
				return errors.New("trying to end a race but is not racing")
			} else if isRacing && (entry.entryType == Start) {
				return errors.New("trying to start a race but is currently racing")
			} else {
				return errors.New("unknown error occured")
			}
		}

		if isRacing {
			return errors.New("finished with no end to last race")
		}
	}

	return nil
}

/*
	Prize Ideas
		Closest Finishes
		Biggest Comeback
		Most Consistent
		Closest to speed of a racoon
		Total time of all races closest to the speed of another world event, like a famous battle.


*/
