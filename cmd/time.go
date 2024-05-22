package cmd

import (
	"github.com/spf13/cobra"
	"github.com/woaitsAryan/regit/internal/scripts"
	"github.com/woaitsAryan/regit/internal/models"
)

var RetimeCommand *cobra.Command
var RewindCommand *cobra.Command
var FastForwardCommand *cobra.Command

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

	RewindCommand = &cobra.Command{
		Use:   "rewind [duration]",
		Short: "Rewinds the commits of a repository by a given duration",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			duration := args[0]
			scripts.Rewind(duration, models.RewindFlags)
		},
	}
	RewindCommand.PersistentFlags().BoolVarP(&models.RewindFlags.Verbose, "verbose", "v", false, "verbose output")
	RewindCommand.PersistentFlags().BoolVarP(&models.RewindFlags.Quiet, "quiet", "q", false, "quiet output")
	RewindCommand.PersistentFlags().StringVarP(&models.RewindFlags.Source, "source", "s", ".", "path to the git repo")
	RewindCommand.PersistentFlags().StringVarP(&models.RewindFlags.Branch, "branch", "b", ".", "specify a branch")

	FastForwardCommand = &cobra.Command{
		Use:   "fastforward [duration]",
		Short: "Fast forwards the commits of a repository by a given duration",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			duration := args[0]
			scripts.FastForward(duration, models.FastForwardFlags)
		},
	}
	FastForwardCommand.PersistentFlags().BoolVarP(&models.FastForwardFlags.Verbose, "verbose", "v", false, "verbose output")
	FastForwardCommand.PersistentFlags().BoolVarP(&models.FastForwardFlags.Quiet, "quiet", "q", false, "quiet output")
	FastForwardCommand.PersistentFlags().StringVarP(&models.FastForwardFlags.Source, "source", "s", ".", "path to the git repo")
	FastForwardCommand.PersistentFlags().StringVarP(&models.FastForwardFlags.Branch, "branch", "b", ".", "specify a branch")
}
