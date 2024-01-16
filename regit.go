package main

import (
	"github.com/spf13/cobra"
	"github.com/woaitsAryan/regit/setup"
)

func main() {
	var rootCmd = &cobra.Command{Use: "regit"}
	rootCmd.AddCommand(setup.OwnGitCommand, setup.BlameGitCommand, setup.BlameLinusCommand)
	cobra.CheckErr(rootCmd.Execute())
}