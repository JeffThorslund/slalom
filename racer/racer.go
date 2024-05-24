package racer

import (
	"fmt"
)

// A person competing the race
type Racer struct {
	Id       RacerId  // unique id given to each racer
	Name     string   // given name of each racer
	Gender   Gender   // gender of each racer
	Category Category // skill level catagory of each racer
}

type RacerId string

type Gender int

const (
	_ Gender = iota
	Male
	Female
)

func (g Gender) String() string {
	switch g {
	case Male:
		return "Male"
	case Female:
		return "Female"
	default:
		return "Unknown"
	}
}

type Category int

const (
	_ Category = iota
	Beginner
	Intermediate
	Advanced
)

func (c Category) String() string {
	switch c {
	case Beginner:
		return "Beginner"
	case Intermediate:
		return "Intermediate"
	case Advanced:
		return "Advanced"
	default:
		return "Unknown"
	}
}

func NewRacer(id RacerId, name string, gender Gender, category Category) Racer {
	return Racer{
		id, name, gender, category,
	}
}

func (r *Racer) String() string {
	return fmt.Sprintf("Racer Name: %s, Category: %v", r.Name, r.Category)
}

type Racers []Racer

type RacersMap map[RacerId]Racer

func (racers *Racers) CreateRacersMap() RacersMap {
	racersMap := make(map[RacerId]Racer)

	for _, racer := range *racers {
		racersMap[racer.Id] = racer
	}

	return racersMap
}
