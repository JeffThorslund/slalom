package main

import (
	"log"
	"time"
)

type entryType int

const (
	_ entryType = iota
	Start
	End
)

// A record of a start or finish, as written by start or end timer
type entry struct {
	racerId   racerId
	time      time.Time
	entryType entryType
}

func newStart(record []string) entry {
	return entry{
		racerId:   racerId(record[0]),
		time:      parseTime(record[1]),
		entryType: Start,
	}
}

func newEnd(record []string) entry {
	return entry{
		racerId:   racerId(record[0]),
		time:      parseTime(record[1]),
		entryType: End,
	}
}

func parseTime(timeStr string) time.Time {
	t, err := time.Parse(time.RFC3339, timeStr)
	if err != nil {
		log.Fatal(err)
	}
	return t
}
