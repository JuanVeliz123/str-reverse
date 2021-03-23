package internal

import (
	"encoding/csv"
	"io"
	"log"
	"os"
)

func ReadCases(path *string) []string {
	f, err := os.Open(*path)
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(f)
	var cases []string
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalln(err)
		}

		cases = append(cases, record[0])
	}
	return cases
}
