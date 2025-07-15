package handler

import (
	"encapsula/database"
	"log"

	"github.com/ewriq/pouch"
)

func Remove(name string) error {
	log.Println("🗑️  Container siliniyor...")
	err = database.Remove(name)
	if err != nil {
		log.Fatalf("❌ Veritabanından silinemedi: %v", err)
	}

	_ ,err := pouch.Remove(name, true)
	if err != nil {
		log.Fatalf("❌ Container silinemedi: %v", err)
		return err
	}

	log.Printf("✅ Container silindi: %s\n", name)
	return nil
}