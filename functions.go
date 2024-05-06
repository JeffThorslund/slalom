package main

import (
	"sort"
	"time"
)

type SortedEntriesPerRacer map[RacerId][]Entry

func createSortedEntriesPerRacer(starts []Entry, ends []Entry) SortedEntriesPerRacer {
	allRaces := append(starts, ends...)

	sort.Slice(allRaces, func(i, j int) bool {
		return allRaces[i].time.Before(allRaces[j].time)
	})

	raceMap := make(SortedEntriesPerRacer)

	for _, race := range allRaces {
		raceMap[race.racerId] = append(raceMap[race.racerId], race)
	}

	return raceMap
}

// A 2 element slice with a start and end entry, representing a completed race
type Race []Entry

func (r *Race) getRaceTime() time.Duration {
	if len(*r) < 2 {
		return 0
	}
	start := (*r)[0].time
	end := (*r)[1].time
	return end.Sub(start)
}

type SortedRacesPerRacer map[RacerId][]Race

func (se SortedEntriesPerRacer) ToRaces() SortedRacesPerRacer {
	sr := make(SortedRacesPerRacer) // Initialize the map

	for racerId, entries := range se {
		var races []Race
		for i := 0; i < len(entries); i += 2 {
			race := entries[i : i+2]
			races = append(races, race)
		}

		sr[racerId] = races
	}

	return sr
}
