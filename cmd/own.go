package cmd

import (
	"github.com/spf13/cobra"
	"github.com/woaitsAryan/regit/internal/scripts"
	"github.com/woaitsAryan/regit/internal/models"
)

var OwnGitCommand *cobra.Command
var BlameGitCommand *cobra.Command
var BlameLinusCommand *cobra.Command

func init() {
	OwnGitCommand = &cobra.Command{
		Use:   "own",
		Short: "Own the commits of a repository",
		Run: func(cmd *cobra.Command, args []string) {
			scripts.Owngit(models.OwnFlags)
		},
	}

	OwnGitCommand.PersistentFlags().BoolVarP(&models.OwnFlags.Verbose, "verbose", "v", false, "verbose output")
	OwnGitCommand.PersistentFlags().BoolVarP(&models.OwnFlags.Quiet, "quiet", "q", false, "quiet output")
	OwnGitCommand.PersistentFlags().StringVarP(&models.OwnFlags.Source, "source", "s", ".", "path to the git repo")
	OwnGitCommand.PersistentFlags().StringVarP(&models.RecommitFlags.Branch, "branch", "b", ".", "specify a branch")

	BlameGitCommand = &cobra.Command{
		Use:   "blame [name] [email]",
		Short: "Blame the commits of a repository",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			name := args[0]
			email := args[1]
			scripts.Blamegit(name, email, models.BlameFlags)
		},
	}
	BlameGitCommand.PersistentFlags().BoolVarP(&models.BlameFlags.Verbose, "verbose", "v", false, "verbose output")
	BlameGitCommand.PersistentFlags().BoolVarP(&models.BlameFlags.Verbose, "quiet", "q", false, "quiet output")
	BlameGitCommand.PersistentFlags().StringVarP(&models.BlameFlags.Source, "source", "s", ".", "path to the git repo")
	BlameGitCommand.PersistentFlags().StringVarP(&models.BlameFlags.Branch, "branch", "b", ".", "specify a branch")

	BlameLinusCommand = &cobra.Command{
		Use:   "blame-linus",
		Short: "Give all your commits to Linus :)",
		Run: func(cmd *cobra.Command, args []string) {
			scripts.Blamegit("torvalds", "torvalds@linux-foundation.org", models.BlameLinusFlags)
		},
	}
	BlameLinusCommand.PersistentFlags().BoolVarP(&models.BlameLinusFlags.Verbose, "verbose", "v", false, "verbose output")
	BlameLinusCommand.PersistentFlags().BoolVarP(&models.BlameLinusFlags.Quiet, "quiet", "q", false, "quiet output")
	BlameLinusCommand.PersistentFlags().StringVarP(&models.BlameLinusFlags.Source, "source", "s", ".", "path to the git repo")
	BlameLinusCommand.PersistentFlags().StringVarP(&models.BlameLinusFlags.Branch, "branch", "b", ".", "specify a branch")
}