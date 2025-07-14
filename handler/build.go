package handler

import (
	"log"

	"github.com/ewriq/pouch"
)

var err error

func Build(name string) string {
	log.Println("🚀 Web SSH (ttyd) container başlatılıyor...")

	opt := pouch.CreateOptions{
		Name:  name,
		Image: "tsl0922/ttyd", 
		Port:  "7681:7681",   
	}

	id, err := pouch.Create(opt)
	if err != nil {
		log.Fatalf("❌ Container oluşturulamadı: %v", err)
	}
	log.Printf("✅ Container oluşturuldu: %s\n", id)


	err = pouch.Start(id)

	if err != nil {
		log.Fatalf("❌ Container başlatılamadı: %v", err,)
	}
	log.Println("✅ Container başlatıldı.")
	log.Println("🌐 Tarayıcıdan aç: http://localhost:7681")
    return id
}

