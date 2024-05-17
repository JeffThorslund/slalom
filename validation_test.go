package main

import (
	"errors"
	"fmt"
	"testing"
)

func TestAssertManyValidRacersInformation(t *testing.T) {
	racers := []racer{
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
		racer racer
		want  error
	}{
		{
			racer{
				id:   "id",
				name: "Jeff",
			},
			nil,
		},
		{
			racer{
				id:   "id",
				name: "",
			},
			ErrEmptyRacerName,
		},
		{
			racer{
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
