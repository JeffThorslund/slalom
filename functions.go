package main

import "sort"

type RaceSummaryPerRacer map[RacerId][]Entry

func createRaceSummaryPerRacer(starts []Entry, ends []Entry) RaceSummaryPerRacer {
	allRaces := append(starts, ends...)

	sort.Slice(allRaces, func(i, j int) bool {
		return allRaces[i].time.Before(allRaces[j].time)
	})

	raceMap := make(RaceSummaryPerRacer)

	for _, race := range allRaces {
		raceMap[race.racerId] = append(raceMap[race.racerId], race)
	}

	return raceMap
}
