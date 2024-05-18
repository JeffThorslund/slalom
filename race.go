package main

import (
	"encoding/csv"
	"math"
	"time"
)

// A 2 element slice with a start and end entry, representing a completed race
type Race struct {
	racerId racerId
	start   time.Time
	end     time.Time
}

func newRace() Race {
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

func (r *Race) getSpeedDiffSeconds(comparator float64) float64 {
	return math.Abs(r.getRaceSpeed() - comparator)
}

func (r *Race) formatRace() []string {
	return []string{
		string((*r).racerId),
		(*r).start.String(),
		(*r).end.String(),
		(*r).getRaceTime().String(),
	}
}

func (r *Race) getHeaders() []string {
	return []string{"id", "start", "end", "total"}
}

type Races []Race

func (rs Races) write(title string, w *csv.Writer) error {

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
