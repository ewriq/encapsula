package database

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func Create(options, name, id string) error {
	const filePath = "./database/db.csv"

	file, err := os.OpenFile(filePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("Dosya açılamadı: %v", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	record := []string{id, name, options}
	if err := writer.Write(record); err != nil {
		return fmt.Errorf("Satır yazılamadı: %v", err)
	}

	log.Println("✅ Kayıt başarıyla eklendi.")
	return nil
}