package main

import (
	"github.com/spf13/cobra"
	"github.com/woaitsAryan/regit/tools"
)

func main() {
	var ownGitCommand = &cobra.Command{
		Use:   "own [path]",
		Short: "Own the commits of a repository",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			tools.Owngit(path)
		},
	}

	var blameGitCommand = &cobra.Command{
		Use:   "blame [path] [name] [email]",
		Short: "Blame the commits of a repository",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			name := args[1]
			email := args[2]
			tools.Blamegit(path, name, email)
		},
	}

	var rootCmd = &cobra.Command{Use: "regit"}
	rootCmd.AddCommand(ownGitCommand, blameGitCommand)
	cobra.CheckErr(rootCmd.Execute())
}