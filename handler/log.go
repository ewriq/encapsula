package handler

import (
	"log"

	"github.com/ewriq/pouch"
)

func LogStats(name string) error {
	stats, err := pouch.ContainerStats(name)
	if err != nil {
		return err
	} else {
		log.Println("Konteyner Kullanım Bilgileri:")
		for k, v := range stats {
			log.Printf("%s: %s\n", k, v)
		}
	}

	return nil
}