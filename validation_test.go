package main

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

func TestAssertManyValidRacersInformation(t *testing.T) {
	racers := []Racer{
		{
			id:   "id1",
			name: "Jeff",
		},
		{
			id:   "id2",
			name: "Bob",
		},
		{
			id:   "id3",
			name: "",
		},
	}

	err := assertManyValidRacersInformation(racers)

	if err == nil {
		t.Fatalf("did not detect error, %v", ErrEmptyRacerName)
	}
}

func TestAssertSingleValidRacerInformation(t *testing.T) {
	var tests = []struct {
		racer Racer
		want  error
	}{
		{
			Racer{
				id:   "id",
				name: "Jeff",
			},
			nil,
		},
		{
			Racer{
				id:   "id",
				name: "",
			},
			ErrEmptyRacerName,
		},
		{
			Racer{
				id:   "",
				name: "Jeff",
			},
			ErrEmptyRacerId,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf("%s:%s", tt.racer.id, tt.racer.name)
		t.Run(testname, func(t *testing.T) {
			err := assertSingleValidRacerInformation(tt.racer)
			if !errors.Is(err, tt.want) {
				t.Errorf("got %v, want %v", err, tt.want)
			}
		})
	}
}

func TestAssertOrderedRaceStarts(t *testing.T) {

	var tests = []struct {
		name  string
		times []time.Time
		want  error
	}{
		{
			name: "ordered",
			times: []time.Time{
				time.Now(),
				time.Now().AddDate(0, 0, 3),
				time.Now().AddDate(0, 0, 6),
			},
			want: nil,
		},
		{
			name: "unordered",
			times: []time.Time{
				time.Now().AddDate(0, 0, 3),
				time.Now().AddDate(0, 0, 5),
				time.Now(),
			},
			want: ErrUnorderedStarts,
		},
	}

	for _, tt := range tests {
		testname := fmt.Sprintf(tt.name)
		t.Run(testname, func(t *testing.T) {
			t.Log(tt.times)
			err := assertOrderedRaceStarts(tt.times)
			if !errors.Is(err, tt.want) {
				t.Errorf("got %v, want %v", err, tt.want)
			}
		})
	}
}
