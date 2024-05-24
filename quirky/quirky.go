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
