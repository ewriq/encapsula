package cli

import (
	"encapsula/handler"
	"log"

	"github.com/spf13/cobra"
)

var remove = &cobra.Command{
	Use:   "remove [name]",
	Short: "You can remove a container",
	Long:  `Where you can control with removing a container`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			log.Println("ERROR: belirsiz argüman")
			return
		}

		argsr := args[0]

		if argsr != "" {
			err := handler.Remove(argsr)
			if err != nil {
				log.Println("ERROR: Container silinemedi")
				return
			}
			return
		} else {
			log.Println("ERROR: Invalid argument")
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(remove)
}
