package main

import (
	"fmt"
	"log"
)

// A person competing the race
type racer struct {
	id       racerId  // unique id given to each racer
	name     string   // given name of each racer
	gender   gender   // gender of each racer
	category category // skill level catagory of each racer
}

type racerId string

type gender int

const (
	_ gender = iota
	Male
	Female
)

func newRacer(record []string) racer {
	return racer{
		id:       racerId(record[0]),
		name:     record[1],
		gender:   parseGender(record[2]),
		category: parseCategory(record[3]),
	}
}

func parseGender(s string) gender {
	var gender gender
	var err error

	switch s {
	case "m":
		gender, err = Male, nil
	case "f":
		gender, err = Female, nil
	default:
		gender, err = 0, fmt.Errorf("invalid gender: %s", s)
	}

	if err != nil {
		log.Fatal(err)
	}

	return gender
}

type category int

const (
	_ category = iota
	Beginner
	Intermediate
	Advanced
)

func parseCategory(s string) category {
	var category category
	var err error

	switch s {
	case "b":
		category, err = Beginner, nil
	case "i":
		category, err = Intermediate, nil
	case "a":
		category, err = Advanced, nil
	default:
		category, err = 0, fmt.Errorf("invalid catergory: %s", s)
	}

	if err != nil {
		log.Fatal(err)
	}

	return category
}

func (r *racer) String() string {
	return fmt.Sprintf("Racer Name: %s, Category: %v", r.name, r.category)
}
