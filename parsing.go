package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func parseCsvData[D Racer | Entry](path string, parser func([]string) D) []D {
	file, err := os.Open(path)

	if err != nil {
		log.Fatal("Error while reading file", err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil {
		fmt.Println("Error reading records")
	}

	var results []D

	for _, record := range records[1:] {
		results = append(results, parser(record))
	}

	return results
}
