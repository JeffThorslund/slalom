package entry

import (
	"time"

	"github.com/JeffThorslund/slalom-results/racer"
)

type EntryType int

const (
	_ EntryType = iota
	Start
	End
)

// A record of a start or finish, as written by start or end timer
type Entry struct {
	RacerId   racer.RacerId
	Time      time.Time
	EntryType EntryType
}

func NewEntry(racerId racer.RacerId, time time.Time, entryType EntryType) Entry {
	return Entry{
		racerId,
		time,
		entryType,
	}
}
