package cli
import (
	"log"

	"github.com/spf13/cobra"
)

var install = &cobra.Command{
	Use:   "install",
	Short: "You can check build",
	Long:  `Where you can control with operations such as opening and closing build`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("Install command")
	
	},
}

func init() {
	rootCmd.AddCommand(install)
}