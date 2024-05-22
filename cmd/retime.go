package cmd

import (
	"github.com/spf13/cobra"
	"github.com/woaitsAryan/regit/internal/scripts"
	"github.com/woaitsAryan/regit/internal/models"
)

var RetimeCommand *cobra.Command

func init() {
	RetimeCommand = &cobra.Command{
		Use:   "retime [duration]",
		Short: "Retime the commits of a repository to a given duration",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			duration := args[0]
			scripts.Retimegit(duration, models.RetimeFlags)
		},
	}
	RetimeCommand.PersistentFlags().BoolVarP(&models.RetimeFlags.Verbose, "verbose", "v", false, "verbose output")
	RetimeCommand.PersistentFlags().BoolVarP(&models.RetimeFlags.Quiet, "quiet", "q", false, "quiet output")
	RetimeCommand.PersistentFlags().StringVarP(&models.RetimeFlags.Source, "source", "s", ".", "path to the git repo")
	RetimeCommand.PersistentFlags().StringVarP(&models.RetimeFlags.Branch, "branch", "b", ".", "specify a branch")
}
