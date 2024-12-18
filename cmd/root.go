package cmd

import (
	"fmt"
	"os"

	"github.com/jayakrishnanMurali/kit/pkg/commands"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kit",
	Short: "Kit is a minimal VCS implemented in GoLang.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Welcome to Kit!!\n Use 'kit --help' to see available commands.")
	},
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Kit",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Kit v0.1")
	},
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create an empty Git repository or reinitialize an existing one",
	Run: func(cmd *cobra.Command, args []string) {
		commands.InitCmd(args)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(initCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
