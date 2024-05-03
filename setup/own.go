package setup

import (
	"github.com/spf13/cobra"
	"github.com/woaitsAryan/regit/scripts"
)

var flags = make(map[string]bool)
var verbose bool
var quiet bool

var OwnGitCommand *cobra.Command
var BlameGitCommand *cobra.Command
var BlameLinusCommand *cobra.Command

func init() {
	OwnGitCommand = &cobra.Command{
		Use:   "own [path]",
		Short: "Own the commits of a repository",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			flags["verbose"] = verbose
			flags["quiet"] = quiet
			scripts.Owngit(path, flags)
		},
	}

	OwnGitCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	OwnGitCommand.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "quiet output")

	BlameGitCommand = &cobra.Command{
		Use:   "blame [path] [name] [email]",
		Short: "Blame the commits of a repository",
		Args:  cobra.ExactArgs(3),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			name := args[1]
			email := args[2]
			flags["verbose"] = verbose
			flags["quiet"] = quiet
			scripts.Blamegit(path, name, email, flags)
		},
	}
	BlameGitCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	BlameGitCommand.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "quiet output")

	BlameLinusCommand = &cobra.Command{
		Use:   "blame-linus [path]",
		Short: "Give all your commits to Linus :)",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			flags["verbose"] = verbose
			flags["quiet"] = quiet
			scripts.Blamegit(path, "torvalds", "torvalds@linux-foundation.org", flags)
		},
	}
	BlameLinusCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	BlameLinusCommand.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "quiet output")
}