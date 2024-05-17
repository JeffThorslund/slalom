package main

import (
	"sort"
	"time"
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

// A 2 element slice with a start and end entry, representing a completed race
type Race struct {
	racerId racerId
	start   time.Time
	end     time.Time
}

func (r *Race) getRaceTime() time.Duration {
	return (*r).end.Sub((*r).start)
}

func (r *Race) printRace() []string {
	return []string{
		string((*r).racerId),
		(*r).start.String(),
		(*r).end.String(),
		(*r).getRaceTime().String(),
	}
}

type SortedRacesPerRacer map[racerId][]Race

func (se SortedEntriesPerRacer) ToRaces() SortedRacesPerRacer {
	sr := make(SortedRacesPerRacer) // Initialize the map

	for racerId, entries := range se {
		var races []Race
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
