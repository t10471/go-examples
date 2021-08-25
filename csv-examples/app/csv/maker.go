package csv

import (
	"bytes"
	_ "embed"
	"encoding/csv"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

//go:embed MOCK_DATA.csv
var mockCSVBytes []byte

func MakerMain(line int, name string) error {
	header, records, err := readBaseRecords()
	if err != nil {
		return err
	}

	f, err := os.Create(name)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	if err := w.Write(header); err != nil {
		return err
	}

	i := 0
	for i < line {
		log.Printf("i: %d", i)
		records = shuffle(records)
		for _, record := range records {
			i += 1
			record[0] = strconv.Itoa(i)
			if err := w.Write(record); err != nil {
				return err
			}
			if i >= line {
				break
			}
		}
	}
	w.Flush()

	if err := w.Error(); err != nil {
		return err
	}
	return nil
}

func readBaseRecords() ([]string, [][]string, error) {
	r := csv.NewReader(bytes.NewBuffer(mockCSVBytes))
	records, err := r.ReadAll()
	if err != nil {
		return nil, nil, err
	}
	return records[0], records[1:], nil
}

func shuffle(records [][]string) [][]string {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(records), func(i, j int) { records[i], records[j] = records[j], records[i] })
	return records
}