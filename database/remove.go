package database

import (
	"encoding/csv"
	"fmt"
	"os"
)

func Remove(name string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("File read error: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("CSV read error: %v", err)
	}
	file.Close()
	var filtered [][]string
	for _, record := range records {
		if len(record) > 1 && record[1] == name {
			continue
		}
		filtered = append(filtered, record)
	}


	file, err = os.Create(filePath)
	if err != nil {
		return fmt.Errorf("File open error: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	if err := writer.WriteAll(filtered); err != nil {
		return fmt.Errorf("Write error: %v", err)
	}

	return nil
}
