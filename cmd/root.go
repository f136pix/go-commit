/*
Copyright © 2024 FILIPE FURLAN <Filipecocinel@gmail.com>
*/
package cmd

import (
	"go-commit/cmd/commit"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "go-commit",
	Short: "An AI based commit message generator",
	Long:  `Go-commit is a CLI tool that generates commit for your application basing itself on the context of it.`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

// addSubcommandPalletes adds all subcommands to the root command
func addSubcommandPalletes() {
	rootCmd.AddCommand(commit.CommitCmd)
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.go-commit.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addSubcommandPalletes()
}
