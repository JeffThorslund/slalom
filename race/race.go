package race

import (
	"encoding/csv"
	"math"
	"sort"
	"time"

	"github.com/JeffThorslund/slalom-results/entry"
	"github.com/JeffThorslund/slalom-results/racer"
)

// A 2 element slice with a start and end entry, representing a completed Race
type Race struct {
	racerId racer.RacerId
	start   time.Time
	end     time.Time
	racer.Racer
}

func NewRace() Race {
	return Race{
		racerId: "0",
		start:   time.Now(),
		end:     time.Now(),
	}
}

func (r *Race) getRaceTime() time.Duration {
	return (*r).end.Sub((*r).start)
}

const raceLengthInMeters = 200.0

// get race speed in meters per second
func (r *Race) getRaceSpeed() float64 {
	return raceLengthInMeters / r.getRaceTime().Seconds()
}

func (r *Race) GetSpeedDiffSeconds(comparator float64) float64 {
	return math.Abs(r.getRaceSpeed() - comparator)
}

const timeFormatString = "15:04:05"

func (r *Race) formatRace() []string {
	return []string{
		string((*r).racerId),
		(*r).start.Format(timeFormatString),
		(*r).end.Format(timeFormatString),
		(*r).getRaceTime().String(),
		(*r).Racer.Name,
		(*r).Racer.Category.String(),
		(*r).Racer.Gender.String(),
	}
}

func (r *Race) getHeaders() []string {
	return []string{"id", "start", "end", "total", "name", "category", "gender"}
}

type Races []Race

func (rs Races) Write(title string, w *csv.Writer) error {

	if len(rs) == 0 {
		return nil
	}

	if err := w.Write([]string{title}); err != nil {
		return err
	}

	headers := rs[0].getHeaders()
	if err := w.Write(headers); err != nil {
		return err
	}

	if err := w.WriteAll(rs.formatRaces()); err != nil {
		return err
	}

	// Write an empty record to add a newline
	if err := w.Write([]string{}); err != nil {
		return err
	}

	return nil
}

func (rs Races) formatRaces() (formattedRaces [][]string) {
	for _, r := range rs {
		formattedRaces = append(formattedRaces, r.formatRace())
	}

	return
}

type sortedRacesPerRacer map[racer.RacerId]Races

func CreateRacesPerRacer(starts []entry.Entry, ends []entry.Entry, racers racer.Racers) sortedRacesPerRacer {

	se := createSortedEntriesPerRacer(starts, ends)

	sr := make(sortedRacesPerRacer)

	racersMap := racers.CreateRacersMap()

	for racerId, entries := range se {
		var races Races
		for i := 0; i < len(entries); i += 2 {

			races = append(races, Race{
				racerId: entries[i].RacerId,
				start:   entries[i].Time,
				end:     entries[i+1].Time,
				Racer:   racersMap[entries[i].RacerId],
			})
		}

		sr[racerId] = races
	}

	return sr
}

// This function separates concerns in a "per racer" context

type sortedEntriesPerRacer map[racer.RacerId][]entry.Entry

func createSortedEntriesPerRacer(starts []entry.Entry, ends []entry.Entry) sortedEntriesPerRacer {
	allRaces := append(starts, ends...)

	sort.Slice(allRaces, func(i, j int) bool {
		return allRaces[i].Time.Before(allRaces[j].Time)
	})

	raceMap := make(sortedEntriesPerRacer)

	for _, race := range allRaces {
		raceMap[race.RacerId] = append(raceMap[race.RacerId], race)
	}

	return raceMap
}

func (sr sortedRacesPerRacer) Write(title string, w *csv.Writer) error {
	if err := w.Write([]string{title}); err != nil {
		return err
	}

	for racerId, races := range sr {
		if err := races.Write(string(racerId), w); err != nil {
			return err
		}
	}

	return nil
}

func (sr sortedRacesPerRacer) Flatten() (allRaces Races) {
	for _, races := range sr {
		allRaces = append(allRaces, races...)
	}

	sort.Slice(allRaces, func(i, j int) bool {
		return allRaces[i].getRaceTime() < allRaces[j].getRaceTime()
	})
	return
}

type CategoryGenderKey struct {
	category racer.Category
	gender   racer.Gender
}

func (races *Races) CreateCatagorizedRaces() []Races {
	categorizedRaces := make(map[CategoryGenderKey]Races)

	for _, race := range *races {
		key := CategoryGenderKey{race.Racer.Category, race.Racer.Gender}
		categorizedRaces[key] = append(categorizedRaces[key], race)
	}

	return []Races{
		categorizedRaces[CategoryGenderKey{racer.Intermediate, racer.Male}],
		categorizedRaces[CategoryGenderKey{racer.Intermediate, racer.Female}],
		categorizedRaces[CategoryGenderKey{racer.Advanced, racer.Male}],
		categorizedRaces[CategoryGenderKey{racer.Advanced, racer.Female}],
	}

}
