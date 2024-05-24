package quirky

import "github.com/JeffThorslund/slalom-results/race"

func ClosetToPenguinSpeed(races race.Races) race.Race {
	speedOfPenguin := 2.0
	clostRace := race.NewRace()
	minDiff := clostRace.GetSpeedDiffSeconds(speedOfPenguin)

	for _, race := range races {
		diff := race.GetSpeedDiffSeconds(speedOfPenguin)
		if diff < minDiff {
			minDiff = diff
			clostRace = race
		}
	}

	return clostRace
}
