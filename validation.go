package main

import (
	"errors"
	"strconv"
	"time"
)

func getValidationError(starts []Entry, ends []Entry, racers []Racer) error {
	return validationErrorAggregator(
		assertValidRacerInformation(racers),
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
	seenIds := make(map[RacerId]bool)
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
