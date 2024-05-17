package main

import (
	"fmt"
	"log"
)

// A person competing the race
type Racer struct {
	id       racerId  // unique id given to each racer
	name     string   // given name of each racer
	gender   Gender   // gender of each racer
	category Category // skill level catagory of each racer
}

type racerId string

type Gender int

const (
	_ Gender = iota
	Male
	Female
)

func createRacer(record []string) Racer {
	return Racer{
		id:       racerId(record[0]),
		name:     record[1],
		gender:   parseGender(record[2]),
		category: parseCategory(record[3]),
	}
}

func parseGender(s string) Gender {
	var gender Gender
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

type Category int

const (
	_ Category = iota
	Beginner
	Intermediate
	Advanced
)

func parseCategory(s string) Category {
	var category Category
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
