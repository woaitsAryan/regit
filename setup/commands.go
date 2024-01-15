package setup

import (
	"github.com/spf13/cobra"
	"github.com/woaitsAryan/regit/tools"
)

var flags = make(map[string]bool)
var verbose bool
var quiet bool

var OwnGitCommand *cobra.Command
var BlameGitCommand *cobra.Command

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
}