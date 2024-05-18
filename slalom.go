package main

import (
	"encoding/csv"
	"log"
	"os"
)

/**
Assumptions: The data is in order of when it happened (validate for this).
*/

func main() {

	// parse the data from csv into arrays
	starts := parseCsvData("testdata/starts.csv", newStart)
	ends := parseCsvData("testdata/ends.csv", newEnd)
	racers := parseCsvData("testdata/racers.csv", newRacer)

	// validate the data without mutation, throw errors for humans to fix
	validationError := getValidationError(starts, ends, racers)

	if validationError != nil {
		log.Fatal(validationError)
	}

	file, err := os.Create("demo.csv")

	if err != nil {
		log.Fatal("Error creating file")
	}

	w := csv.NewWriter(file)

	// now we know (hopefully) the results are valid, we want to see the results in different views (think selectors)
	// we can continue without validation

	// 1. a "per person" breakdown of their races, sorted. This is the most natural way of constructing the structure so we start with that.
	sortedRacesPerRacer := createSortedEntriesPerRacer(starts, ends).ToRaces()
	if err := sortedRacesPerRacer.write("sorted racers", w); err != nil {
		log.Fatalln(err)
	}

	allRaces := sortedRacesPerRacer.flatten()
	if err := allRaces.write("All Races", w); err != nil {
		log.Fatalln("Error writing all races", err)
	}

	racersMap := make(map[racerId]racer)

	for _, racer := range racers {
		racersMap[racer.id] = racer
	}

	type CategoryGenderKey struct {
		category category
		gender   gender
	}

	categorizedRaces := make(map[CategoryGenderKey]Races)

	for _, race := range allRaces {
		racer := racersMap[race.racerId]
		log.Println(racer.String())
		key := CategoryGenderKey{racer.category, racer.gender}
		categorizedRaces[key] = append(categorizedRaces[key], race)
	}

	categorizedRaces[CategoryGenderKey{Intermediate, Male}].write("IM", w)
	categorizedRaces[CategoryGenderKey{Intermediate, Female}].write("IF", w)
	categorizedRaces[CategoryGenderKey{Advanced, Male}].write("AM", w)
	categorizedRaces[CategoryGenderKey{Advanced, Female}].write("AF", w)

	// 3. the "fun awards"

	c := closetToPenguinSpeed(allRaces)
	log.Println("race", c.formatRace())

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

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
