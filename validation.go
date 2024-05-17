package main

import (
	"errors"
	"fmt"
	"strconv"
	"time"
)

func getValidationError(starts []entry, ends []entry, racers []Racer) error {
	return validationErrorAggregator(
		assertManyValidRacersInformation(racers),
		assertNoDuplicateRacers(racers),
		assertOrderedRaceStarts(getTimes(starts)),
		assertThatAllRacesEnd(starts, ends),
	)
}

// return the first error encountered
func validationErrorAggregator(validationErrors ...error) error {
	for _, validationError := range validationErrors {
		if validationError != nil {
			return validationError
		}
	}

	return nil
}

func assertManyValidRacersInformation(racers []Racer) error {
	for i, racer := range racers {
		err := assertSingleValidRacerInformation(racer)
		if err != nil {
			return fmt.Errorf("index: %d, %w", i, err)
		}
	}

	return nil
}

var ErrEmptyRacerId = errors.New("empty racer id")
var ErrEmptyRacerName = errors.New("empty racer name")

func assertSingleValidRacerInformation(racer Racer) error {
	if racer.id == "" {
		return ErrEmptyRacerId
	}

	if racer.name == "" {
		return ErrEmptyRacerName
	}

	return nil
}

func getTimes(starts []entry) []time.Time {
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
	seenIds := make(map[racerId]bool)
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

func assertThatAllRacesEnd(starts []entry, ends []entry) error {

	raceSummaryPerRacer := createSortedEntriesPerRacer(starts, ends)

	for _, entries := range raceSummaryPerRacer {

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
