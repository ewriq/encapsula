package database

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func List() error {

	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("Dosya açılamadı: %v", err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("CSV verisi okunamadı: %v", err)
	}

	if len(records) == 0 {
		log.Println("📂 Hiç kayıt yok.")
		return nil
	}

	log.Println("📋 Kayıt Listesi:")
	log.Println("ID\tName\tOptions")
	for _, record := range records {
		if len(record) >= 3 {
			log.Printf("%s\t%s\t%s\n", record[0], record[1], record[2])
		} else {
			log.Printf("%v\n", record) 
		}
	}

	return nil
}
