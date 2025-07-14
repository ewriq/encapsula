package cli

import (
	"encapsula/database"
	"log"

	"github.com/spf13/cobra"
)

var build = &cobra.Command{
	Use:   "build",
	Short: "You can check build",
	Long:  `Where you can control with operations such as opening and closing build`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("ERROR: Undefiend arguments")
			return
		}

		argsr := args[0]
		
		if argsr == "ubuntu" || argsr == "debian" || argsr == "arch" || argsr == "fedora" {
			log.Println(argsr)
			ERR := database.Create(argsr, "azazazaza", "1234567890")
			if ERR != nil {
				log.Printf("ERROR: %v", ERR)
			} else {
				log.Println("✅ Build operation completed successfully.")
			}

		} else {
			log.Println("ERROR: Invalid argument")
	}
	},
}

func init() {
	rootCmd.AddCommand(build)
}