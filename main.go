package main

import (
	"github.com/spf13/cobra"
	"github.com/woaitsAryan/regit/internal/initializers"
	"github.com/woaitsAryan/regit/cmd"
)

var rootCmd = &cobra.Command{
	Use:   "regit",
	Short: "Regit is a CLI for managing git repositories",
	Version: initializers.GetLatestTag(),
}

func init() {
	initializers.CheckCommand()
}

func main() {
	rootCmd.AddCommand(cmd.OwnGitCommand)
	rootCmd.AddCommand(cmd.BlameGitCommand)
	rootCmd.AddCommand(cmd.BlameLinusCommand)
	rootCmd.AddCommand(cmd.RetimeCommand)
	rootCmd.AddCommand(cmd.NukeGitCommand)
	rootCmd.AddCommand(cmd.RecommitGitCommand)
	cobra.CheckErr(rootCmd.Execute())
}
