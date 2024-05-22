package cmd

import (
	"github.com/spf13/cobra"
	"github.com/woaitsAryan/regit/internal/scripts"
	"github.com/woaitsAryan/regit/internal/models"
)

var RecommitGitCommand *cobra.Command

func init() {
	RecommitGitCommand = &cobra.Command{
		Use:   "recommit",
		Short: "Rewrite the commit messages of a repo with better messages",
		Run: func(cmd *cobra.Command, args []string) {
			scripts.Recommitgit(models.RecommitFlags)
		},
	}

	RecommitGitCommand.PersistentFlags().BoolVarP(&models.RecommitFlags.Verbose, "verbose", "v", false, "verbose output")
	RecommitGitCommand.PersistentFlags().BoolVarP(&models.RecommitFlags.Quiet, "quiet", "q", false, "quiet output")
	RecommitGitCommand.PersistentFlags().StringVarP(&models.RecommitFlags.Source, "source", "s", ".", "path to the git repo")
	RecommitGitCommand.PersistentFlags().StringVarP(&models.RecommitFlags.Branch, "branch", "b", ".", "specify a branch")
}