package setup

import (
	"github.com/spf13/cobra"
	"github.com/woaitsAryan/regit/scripts"
)

var RecommitGitCommand *cobra.Command

func init() {
	RecommitGitCommand = &cobra.Command{
		Use:   "recommit [path]",
		Short: "Rewrite the commit messages of a repo with better messages",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			flags["verbose"] = verbose
			flags["quiet"] = quiet
			scripts.Recommitgit(path, flags)
		},
	}

	RecommitGitCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	RecommitGitCommand.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "quiet output")
}