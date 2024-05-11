package cmd

import (
	"github.com/spf13/cobra"
	"github.com/woaitsAryan/regit/internal/scripts"
	"github.com/woaitsAryan/regit/internal/models"

)

var NukeGitCommand *cobra.Command

func init() {
	NukeGitCommand = &cobra.Command{
		Use:   "nuke [file]",
		Short: "Nuke the file from all the commits.",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			file := args[0]
			scripts.NukeGit(file, models.NukeFlags)
		},
	}

	NukeGitCommand.PersistentFlags().BoolVarP(&models.NukeFlags.Verbose, "verbose", "v", false, "verbose output")
	NukeGitCommand.PersistentFlags().BoolVarP(&models.NukeFlags.Verbose, "quiet", "q", false, "quiet output")
	NukeGitCommand.PersistentFlags().StringVarP(&models.NukeFlags.Source, "source", "s", ".", "path to the git repo")
}
