package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"sort"
)

/**
Assumptions: The data is in order of when it happened (validate for this).
*/

func main() {
	// parse the data from csv into arrays
	starts := parseCsvData("testdata/starts.csv", createStart)
	ends := parseCsvData("testdata/ends.csv", createEnd)
	racers := parseCsvData("testdata/racers.csv", createRacer)

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

	// Filter this by catagory when catagories are added to racers.

	// 3. the "fun awards"
	/*
		- closest to running pace of a penguin (1.6 miles an hour).
		- most improved
		- most races
		- closest time on the water to a soft boiled egg (4 min)
		- closest race time to attention span of a feral pig
		- money bags award, if time was money, this racer would be the richest (check first)
		- How long does it take to pee 1.5 liters? gold painted toilet seat, and a container of apple juice. Urine Luck.

		-
	*/

	// Deal with this last.

	// finally as a last step, we create a csv file(s) of the results
	// Per racer info sheets, list of all races, list of catagorized races

	file, err := os.Create("demo.csv")

	if err != nil {
		log.Fatal("Error creating file")
	}

	w := csv.NewWriter(file)

	w.Write([]string{"racer id", "start time", "end time", "total time"})

	for _, race := range masterRaceList {

		record := race.printRace()

		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}

	fmt.Println(sortedRacesPerRacer, masterRaceList)
}
