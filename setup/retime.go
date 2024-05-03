package setup

import (
	"github.com/spf13/cobra"
	"github.com/woaitsAryan/regit/scripts"
)

var RetimeCommand *cobra.Command

func init() {
	RetimeCommand = &cobra.Command{
		Use:   "retime [path] [duration]",
		Short: "Retime the commits of a repository to a given duration",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			path := args[0]
			duration := args[1]
			flags["verbose"] = verbose
			flags["quiet"] = quiet
			scripts.Retimegit(path, duration, flags)
		},
	}
	RetimeCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	RetimeCommand.PersistentFlags().BoolVarP(&quiet, "quiet", "q", false, "quiet output")
}
