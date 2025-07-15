package cli

import (
	"encapsula/database"
	"log"

	"github.com/spf13/cobra"
)

var list = &cobra.Command{
	Use:   "list",
	Short: "You can check build",
	Long:  `Where you can control with operations such as opening and closing build`,
	Run: func(cmd *cobra.Command, args []string) {
	err := database.List()
	if err != nil {
		log.Println("ERROR: Unable to list containers")
		return
	}
		log.Println("✅ Containers listed successfully.")
	},
}

func init() {
	rootCmd.AddCommand(list)
}