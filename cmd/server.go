package cmd

import (
	"fmt"
	"vervain/internal/server"

	"github.com/spf13/cobra"
)

// serverCmd represents the server command
var serverCmd = &cobra.Command{
	Use:   "serve",
	Short: "Start the Vervain web server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Starting Vervain web server...")
		server.StartServer()
	},
}

func init() {
	rootCmd.AddCommand(serverCmd)
}
