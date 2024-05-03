package setup

import (
	"github.com/spf13/cobra"
	"github.com/woaitsAryan/regit/scripts"
)


var NukeGitCommand *cobra.Command

func init() {
	NukeGitCommand = &cobra.Command{
		Use:   "nuke [path] [file]",
		Short: "Nuke the file from all the commits.",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			file := args[1]
			flags["verbose"] = verbose
			flags["quiet"] = quiet
			scripts.NukeGit(path, file, flags)
		},
	}

	NukeGitCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	NukeGitCommand.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "quiet output")
}