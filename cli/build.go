package cli

import (
	"encapsula/database"
	"encapsula/handler"
	"log"

	"github.com/spf13/cobra"
)

var build = &cobra.Command{
	Use:   "build",
	Short: "You can check build",
	Long:  `Where you can control with operations such as opening and closing build`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("ERROR: belirsiz argüman")
			return
		}

		argsr := args[0]
		
		if argsr != ""  {
			id := handler.Build(argsr)
			if id == "" {
				log.Println("ERROR: Container oluşturulamadı")
				return
				
			}
			log.Println(argsr)
			err := database.Create(	"ubunutu", argsr, id)
			if err != nil {
				log.Println("ERROR: %v", err)
			} else {
				log.Println("✅ Derleme ve indirme başarılı.")
				return
			}

		} else {
			log.Println("ERROR: Invalid argument")
			return
	}
	},
}

func init() {
	rootCmd.AddCommand(build)
}