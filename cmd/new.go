package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var template string
var path string

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:   "new [project-name]",
	Short: "Create a new Vervain project",
	Long:  `Generates a new Vervain project with optional templates and directory paths.`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		projectName := args[0]

		// Check if path is defined; use default otherwise
		if path == "" {
			path = "./" + projectName
		}

		fmt.Printf("Creating project %s at %s with template %s\n", projectName, path, template)

		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			fmt.Println("Error creating directory:", err)
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Flags for the new command
	newCmd.Flags().StringVarP(&template, "template", "t", "web", "Specify project template (api, web, fullstack)")
	newCmd.Flags().StringVarP(&path, "path", "p", "", "Specify directory for the project")
}
