package main

func closetToPenguinSpeed(races Races) Race {
	speedOfPenguin := 2.0
	clostRace := newRace()
	minDiff := clostRace.getSpeedDiffSeconds(speedOfPenguin)

	for _, race := range races {
		diff := race.getSpeedDiffSeconds(speedOfPenguin)
		if diff < minDiff {
			minDiff = diff
			clostRace = race
		}
	}

	return clostRace
}
