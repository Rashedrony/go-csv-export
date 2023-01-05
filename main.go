package main

import (
	"encoding/csv"
	"log"
	"os"
)

func main() {
	records := [][]string{
		{"first_name", "last_name", "occupation"},
		{"John", "Doe", "Software Engineer"},
		{"Rashed", "Rahman", "Developer"},
	}

	f, err := os.Create("user.csv")

	defer f.Close()

	if err != nil {
		log.Fatal("Failed to open file", err)
	}

	writer := csv.NewWriter(f)
	defer writer.Flush()

	for _, record := range records {
		if err := writer.Write(record); err != nil {
			log.Fatal("Error Writing Record", err)
		}
	}

	// another procedure
	f1, err1 := os.Create("user1.csv")

	if err1 != nil {
		log.Fatal("Failed to open file", err)
	}

	writer1 := csv.NewWriter(f1)
	err1 = writer1.WriteAll(records)

	if err1 != nil {
		log.Fatal("Error Writing Record", err)
	}
}
