package main

import (
	"sort"
)

type SortedEntriesPerRacer map[racerId][]entry

func createSortedEntriesPerRacer(starts []entry, ends []entry) SortedEntriesPerRacer {
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

type SortedRacesPerRacer map[racerId]Races

func (se SortedEntriesPerRacer) ToRaces() SortedRacesPerRacer {
	sr := make(SortedRacesPerRacer) // Initialize the map

	for racerId, entries := range se {
		var races Races
		for i := 0; i < len(entries); i += 2 {
			races = append(races, Race{
				racerId: entries[i].racerId,
				start:   entries[i].time,
				end:     entries[i+1].time,
			})
		}

		sr[racerId] = races
	}

	return sr
}
