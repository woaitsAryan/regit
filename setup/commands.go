package setup

import (
	"github.com/spf13/cobra"
	"github.com/woaitsAryan/regit/tools"
)

var flags = make(map[string]bool)
var verbose bool
var quiet bool
// var timezone string

var OwnGitCommand *cobra.Command
var BlameGitCommand *cobra.Command
var BlameLinusCommand *cobra.Command
var RetimeCommand *cobra.Command


func init() {
	OwnGitCommand = &cobra.Command{
		Use:   "own [path]",
		Short: "Own the commits of a repository",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			flags["verbose"] = verbose
			flags["quiet"] = quiet
			tools.Owngit(path, flags)
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
			tools.Blamegit(path, name, email, flags)
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
			tools.Blamegit(path, "torvalds", "torvalds@linux-foundation.org", flags)
		},
	}
	BlameLinusCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	BlameLinusCommand.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "quiet output")

	RetimeCommand = &cobra.Command{
		Use:   "retime [path] [duration]",
		Short: "Retime the commits of a repository to a given duration",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			duration := args[1]
			flags["verbose"] = verbose
			flags["quiet"] = quiet
			tools.Retimegit(path, duration, flags)
		},
	}
	RetimeCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	RetimeCommand.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "quiet output")
	// RetimeCommand.PersistentFlags().StringVarP(&timezone, "timezone", "t", "IST", "timezone to retime to")
}