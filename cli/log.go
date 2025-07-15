package cli

import (
	"encapsula/handler"
	"log"
	"time"

	"github.com/spf13/cobra"
)

var logs = &cobra.Command{
	Use:   "log [option] [name]",
	Short: "You can check log",
	Long:  `Control log operations such as container stats or console output.`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 2 {
			log.Println("❌ HATA: Kullanım: log [stats|console] [container_name]")
			return
		}

		option := args[0]
		name := args[1]

		switch option {
		case "stats":
			for {
				err := handler.LogStats(name)
				if err != nil {
					log.Printf("❌ HATA: %v\n", err)
					return
				}
				log.Println("----------------------------------")
				time.Sleep(10 * time.Second)
			}

		case "console":
			err := handler.LogStats(name)
			if err != nil {
				log.Printf("❌ HATA: %v\n", err)
			}

		default:
			log.Printf("❌ HATA: Geçersiz seçenek: '%s'\n", option)
		}
	},
}

func init() {
	rootCmd.AddCommand(logs)
}
