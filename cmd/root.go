package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kit",
	Short: "Kit is a minimal VCS implemented in GoLang.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Kit!!\n Use 'kit --help' to see available commands.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
