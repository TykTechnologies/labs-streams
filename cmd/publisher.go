package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var publisherCmd = &cobra.Command{
	Use:   "publisher",
	Short: "A brief description of your command",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("publisher called")

	},
}

func init() {
	rootCmd.AddCommand(publisherCmd)
}
